package ai.verta.modeldb.common.artifactStore;

import ai.verta.modeldb.GetUrlForArtifact;
import ai.verta.modeldb.GetUrlForArtifact.Response;
import ai.verta.modeldb.common.CommonConstants;
import ai.verta.modeldb.common.HttpCodeToGRPCCode;
import ai.verta.modeldb.common.artifactStore.storageservice.ArtifactStoreService;
import ai.verta.modeldb.common.config.ArtifactStoreConfig;
import ai.verta.modeldb.common.exceptions.ModelDBException;
import com.amazonaws.AmazonServiceException;
import com.amazonaws.SdkClientException;
import com.amazonaws.services.s3.model.PartETag;
import com.google.rpc.Code;
import java.io.InputStream;
import java.util.List;
import java.util.Optional;
import org.apache.logging.log4j.LogManager;
import org.apache.logging.log4j.Logger;

public class ArtifactStoreDAORdbImpl implements ArtifactStoreDAO {

  private static final Logger LOGGER = LogManager.getLogger(ArtifactStoreDAORdbImpl.class);
  private ArtifactStoreService artifactStoreService;
  private final ArtifactStoreConfig artifactStoreConfig;

  public ArtifactStoreDAORdbImpl(
      ArtifactStoreService artifactStoreService, ArtifactStoreConfig artifactStoreConfig) {
    this.artifactStoreService = artifactStoreService;
    this.artifactStoreConfig = artifactStoreConfig;
  }

  @Override
  public GetUrlForArtifact.Response getUrlForArtifact(String s3Key, String method)
      throws ModelDBException {
    return getUrlForArtifactMultipart(s3Key, method, 0, null);
  }

  @Override
  public Response getUrlForArtifactMultipart(
      String s3Key, String method, long partNumber, String uploadId) throws ModelDBException {
    try {
      String presignedUrl =
          artifactStoreService.generatePresignedUrl(s3Key, method, partNumber, uploadId);
      return GetUrlForArtifact.Response.newBuilder()
          .setMultipartUploadOk(
              artifactStoreConfig.getArtifactStoreType().equals(CommonConstants.S3)
                  && uploadId != null)
          .setUrl(presignedUrl)
          .build();
    } catch (AmazonServiceException e) {
      // Amazon S3 couldn't be contacted for a response, or the client
      // couldn't parse the response from Amazon S3.
      String errorMessage = e.getMessage();
      LOGGER.warn(errorMessage);
      throw new ModelDBException(
          errorMessage, HttpCodeToGRPCCode.convertHTTPCodeToGRPCCode(e.getStatusCode()));
    } catch (SdkClientException ex) {
      String errorMessage = ex.getMessage();
      LOGGER.warn(errorMessage);
      throw new ModelDBException(errorMessage, Code.INTERNAL);
    }
  }

  @Override
  public Optional<String> initializeMultipart(String s3Key) throws ModelDBException {
    return artifactStoreService.initiateMultipart(s3Key);
  }

  @Override
  public void commitMultipart(String s3Path, String uploadId, List<PartETag> partETags)
      throws ModelDBException {
    artifactStoreService.commitMultipart(s3Path, uploadId, partETags);
  }

  @Override
  public InputStream downloadArtifact(String artifactPath) throws ModelDBException {
    return artifactStoreService.downloadFileFromStorage(artifactPath);
  }
}
