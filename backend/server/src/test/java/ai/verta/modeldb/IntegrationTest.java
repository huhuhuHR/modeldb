package ai.verta.modeldb;

import static org.junit.Assert.*;
import static org.mockito.ArgumentMatchers.any;
import static org.mockito.Mockito.when;
import static org.springframework.boot.test.context.SpringBootTest.WebEnvironment.DEFINED_PORT;

import ai.verta.common.Artifact;
import ai.verta.common.ArtifactTypeEnum.ArtifactType;
import ai.verta.common.KeyValue;
import ai.verta.common.ModelDBResourceEnum.ModelDBServiceResourceTypes;
import ai.verta.common.ValueTypeEnum.ValueType;
import ai.verta.uac.Action;
import ai.verta.uac.GetSelfAllowedResources;
import ai.verta.uac.ModelDBActionEnum.ModelDBServiceActions;
import ai.verta.uac.ResourceType;
import ai.verta.uac.ServiceEnum;
import com.google.common.util.concurrent.Futures;
import com.google.protobuf.ListValue;
import com.google.protobuf.Value;
import io.grpc.Status;
import io.grpc.StatusRuntimeException;
import java.util.ArrayList;
import java.util.Calendar;
import java.util.List;
import java.util.Map;
import java.util.Set;
import org.apache.logging.log4j.LogManager;
import org.apache.logging.log4j.Logger;
import org.junit.jupiter.api.AfterEach;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.extension.ExtendWith;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.test.context.ContextConfiguration;
import org.springframework.test.context.junit.jupiter.SpringExtension;

@ExtendWith(SpringExtension.class)
@SpringBootTest(classes = App.class, webEnvironment = DEFINED_PORT)
@ContextConfiguration(classes = {ModeldbTestConfigurationBeans.class})
public class IntegrationTest extends ModeldbTestSetup {

  public static final Logger LOGGER = LogManager.getLogger(IntegrationTest.class);

  private CreateProject createProjectRequest() {
    String projectName = "project_" + Calendar.getInstance().getTimeInMillis();

    List<KeyValue> metadataList = new ArrayList<>();
    for (int count = 0; count < 5; count++) {
      Value stringValue =
          Value.newBuilder()
              .setStringValue(
                  "attribute_" + count + "_" + Calendar.getInstance().getTimeInMillis() + "_value")
              .build();
      KeyValue keyValue =
          KeyValue.newBuilder()
              .setKey("attribute_" + count + "_" + Calendar.getInstance().getTimeInMillis())
              .setValue(stringValue)
              .setValueType(ValueType.STRING)
              .build();
      metadataList.add(keyValue);
    }

    return CreateProject.newBuilder()
        .setName(projectName)
        .setDescription("This is a project description.")
        .addTags("tag_1_" + Calendar.getInstance().getTimeInMillis())
        .addTags("tag_2_" + +Calendar.getInstance().getTimeInMillis())
        .addAllAttributes(metadataList)
        .build();
  }

  private CreateExperiment createExperimentRequest(String projectId) {
    List<KeyValue> attributeList = new ArrayList<>();
    Value stringValue =
        Value.newBuilder()
            .setStringValue("attributes_value_" + Calendar.getInstance().getTimeInMillis())
            .build();
    attributeList.add(
        KeyValue.newBuilder()
            .setKey("attribute_1_" + Calendar.getInstance().getTimeInMillis())
            .setValue(stringValue)
            .setValueType(ValueType.STRING)
            .build());
    stringValue =
        Value.newBuilder()
            .setStringValue("attributes_value_" + Calendar.getInstance().getTimeInMillis())
            .build();
    attributeList.add(
        KeyValue.newBuilder()
            .setKey("attribute_2_" + Calendar.getInstance().getTimeInMillis())
            .setValue(stringValue)
            .setValueType(ValueType.STRING)
            .build());

    return CreateExperiment.newBuilder()
        .setProjectId(projectId)
        .setName("experiment_" + Calendar.getInstance().getTimeInMillis())
        .setDescription("This is a experiment description.")
        .addTags("tag_1" + Calendar.getInstance().getTimeInMillis())
        .addTags("tag_2" + +Calendar.getInstance().getTimeInMillis())
        .addAllAttributes(attributeList)
        .build();
  }

  private CreateExperimentRun createExperimentRunRequest(String projectId, String experimentId) {

    List<String> tags = new ArrayList<>();
    tags.add("Tag 1 " + Calendar.getInstance().getTimeInMillis());
    tags.add("Tag 2 " + Calendar.getInstance().getTimeInMillis());

    List<KeyValue> attributeList = new ArrayList<>();
    Value intValue =
        Value.newBuilder().setNumberValue(Calendar.getInstance().getTimeInMillis()).build();
    attributeList.add(
        KeyValue.newBuilder()
            .setKey("attribute_1_" + Calendar.getInstance().getTimeInMillis())
            .setValue(intValue)
            .setValueType(ValueType.NUMBER)
            .build());
    Value stringValue =
        Value.newBuilder()
            .setStringValue("attributes_value_" + Calendar.getInstance().getTimeInMillis())
            .build();
    attributeList.add(
        KeyValue.newBuilder()
            .setKey("attribute_2_" + Calendar.getInstance().getTimeInMillis())
            .setValue(stringValue)
            .setValueType(ValueType.STRING)
            .build());

    List<KeyValue> hyperparameters = new ArrayList<>();
    intValue =
        Value.newBuilder().setNumberValue(1 + Calendar.getInstance().getTimeInMillis()).build();
    hyperparameters.add(
        KeyValue.newBuilder()
            .setKey("hyperparameters_1_" + Calendar.getInstance().getTimeInMillis())
            .setValue(intValue)
            .setValueType(ValueType.NUMBER)
            .build());
    stringValue =
        Value.newBuilder()
            .setStringValue("hyperparameters_value_" + Calendar.getInstance().getTimeInMillis())
            .build();
    hyperparameters.add(
        KeyValue.newBuilder()
            .setKey("hyperparameters_2_" + Calendar.getInstance().getTimeInMillis())
            .setValue(stringValue)
            .setValueType(ValueType.STRING)
            .build());

    List<Artifact> artifactList = new ArrayList<>();
    artifactList.add(
        Artifact.newBuilder()
            .setKey("Google developer Artifact")
            .setPath(
                "https://www.google.co.in/imgres?imgurl=https%3A%2F%2Flh3.googleusercontent.com%2FFyZA5SbKPJA7Y3XCeb9-uGwow8pugxj77Z1xvs8vFS6EI3FABZDCDtA9ScqzHKjhU8av_Ck95ET-P_rPJCbC2v_OswCN8A%3Ds688&imgrefurl=https%3A%2F%2Fdevelopers.google.com%2F&docid=1MVaWrOPIjYeJM&tbnid=I7xZkRN5m6_z-M%3A&vet=10ahUKEwjr1OiS0ufeAhWNbX0KHXpFAmQQMwhyKAMwAw..i&w=688&h=387&bih=657&biw=1366&q=google&ved=0ahUKEwjr1OiS0ufeAhWNbX0KHXpFAmQQMwhyKAMwAw&iact=mrc&uact=8")
            .setArtifactType(ArtifactType.BLOB)
            .setUploadCompleted(true)
            .build());
    artifactList.add(
        Artifact.newBuilder()
            .setKey("Google Pay Artifact")
            .setPath(
                "https://www.google.co.in/imgres?imgurl=https%3A%2F%2Fpay.google.com%2Fabout%2Fstatic%2Fimages%2Fsocial%2Fknowledge_graph_logo.png&imgrefurl=https%3A%2F%2Fpay.google.com%2Fabout%2F&docid=zmoE9BrSKYr4xM&tbnid=eCL1Y6f9xrPtDM%3A&vet=10ahUKEwjr1OiS0ufeAhWNbX0KHXpFAmQQMwhwKAIwAg..i&w=1200&h=630&bih=657&biw=1366&q=google&ved=0ahUKEwjr1OiS0ufeAhWNbX0KHXpFAmQQMwhwKAIwAg&iact=mrc&uact=8")
            .setArtifactType(ArtifactType.IMAGE)
            .setUploadCompleted(true)
            .build());

    List<Artifact> datasets = new ArrayList<>();
    datasets.add(
        Artifact.newBuilder()
            .setKey("Google developer datasets")
            .setPath("This is data artifact type in Google developer datasets")
            .setArtifactType(ArtifactType.MODEL)
            .setUploadCompleted(true)
            .build());
    datasets.add(
        Artifact.newBuilder()
            .setKey("Google Pay datasets")
            .setPath("This is data artifact type in Google Pay datasets")
            .setArtifactType(ArtifactType.DATA)
            .setUploadCompleted(true)
            .build());

    List<KeyValue> metrics = new ArrayList<>();
    Value listValue =
        Value.newBuilder().setListValue(ListValue.newBuilder().addValues(intValue)).build();
    metrics.add(
        KeyValue.newBuilder()
            .setKey("metrics_1_" + Calendar.getInstance().getTimeInMillis())
            .setValue(listValue)
            .setValueType(ValueType.LIST)
            .build());
    stringValue =
        Value.newBuilder()
            .setStringValue("metrics_value_" + Calendar.getInstance().getTimeInMillis())
            .build();
    metrics.add(
        KeyValue.newBuilder()
            .setKey("metrics_2_" + Calendar.getInstance().getTimeInMillis())
            .setValue(stringValue)
            .setValueType(ValueType.STRING)
            .build());

    List<Observation> observations = new ArrayList<>();
    stringValue =
        Value.newBuilder()
            .setStringValue("observation_value_" + Calendar.getInstance().getTimeInMillis())
            .build();
    observations.add(
        Observation.newBuilder()
            .setAttribute(
                KeyValue.newBuilder()
                    .setKey("Observation Key " + Calendar.getInstance().getTimeInMillis())
                    .setValue(stringValue)
                    .setValueType(ValueType.STRING))
            .setTimestamp(Calendar.getInstance().getTimeInMillis())
            .setEpochNumber(Value.newBuilder().setNumberValue(0.0).build())
            .build());

    List<Feature> features = new ArrayList<>();
    features.add(Feature.newBuilder().setName("ExperimentRun Test case feature 1").build());
    features.add(Feature.newBuilder().setName("ExperimentRun Test case feature 2").build());

    return CreateExperimentRun.newBuilder()
        .setProjectId(projectId)
        .setExperimentId(experimentId)
        .setName("ExperimentRun Name_" + Calendar.getInstance().getTimeInMillis())
        .setDescription("this is a ExperimentRun description")
        .setDateCreated(Calendar.getInstance().getTimeInMillis())
        .setDateUpdated(Calendar.getInstance().getTimeInMillis())
        .setStartTime(Calendar.getInstance().getTime().getTime())
        .setEndTime(Calendar.getInstance().getTime().getTime())
        .setCodeVersion("1.0")
        .addAllTags(tags)
        .addAllAttributes(attributeList)
        .addAllHyperparameters(hyperparameters)
        .addAllArtifacts(artifactList)
        .addAllDatasets(datasets)
        .addAllMetrics(metrics)
        .addAllObservations(observations)
        .addAllFeatures(features)
        .build();
  }

  @BeforeEach
  public void createEntities() {
    initializeChannelBuilderAndExternalServiceStubs();

    if (isRunningIsolated()) {
      setupMockUacEndpoints(uac);
    }
  }

  @AfterEach
  public void removeEntities() {
    cleanUpResources();
  }

  @Test
  public void a_AllEntityCRUDTest() {
    LOGGER.info("All Entity CRUD Test start................................");
    try {
      LOGGER.info("Create Project test start................................");
      CreateProject createProjectRequest = createProjectRequest();
      CreateProject.Response createProjectRequestResponse =
          projectServiceStub.createProject(createProjectRequest);
      Project project = createProjectRequestResponse.getProject();
      assertEquals(
          "Project name not match with expected project name",
          createProjectRequest.getName(),
          project.getName());
      LOGGER.info("Create Project test successfully executed");
      LOGGER.info("Create Project test stop................................");

      if (isRunningIsolated()) {
        mockGetResourcesForAllProjects(Map.of(project.getId(), project), testUser1);
        when(uac.getAuthzService()
                .getSelfAllowedResources(
                    GetSelfAllowedResources.newBuilder()
                        .addActions(
                            Action.newBuilder()
                                .setModeldbServiceAction(ModelDBServiceActions.READ)
                                .setService(ServiceEnum.Service.MODELDB_SERVICE))
                        .setService(ServiceEnum.Service.MODELDB_SERVICE)
                        .setResourceType(
                            ResourceType.newBuilder()
                                .setModeldbServiceResourceType(
                                    ModelDBServiceResourceTypes.REPOSITORY))
                        .build()))
            .thenReturn(
                Futures.immediateFuture(GetSelfAllowedResources.Response.newBuilder().build()));
      }

      LOGGER.info("Create Experiment test start................................");
      CreateExperiment createExperimentRequest = createExperimentRequest(project.getId());
      CreateExperiment.Response createExperimentResponse =
          experimentServiceStub.createExperiment(createExperimentRequest);
      Experiment experiment = createExperimentResponse.getExperiment();
      assertEquals(
          "Experiment name not match with expected Experiment name",
          createExperimentRequest.getName(),
          experiment.getName());
      LOGGER.info("Create Experiment test successfully executed");

      CreateExperiment createExperimentRequest2 = createExperimentRequest(project.getId());
      CreateExperiment.Response createExperimentResponse2 =
          experimentServiceStub.createExperiment(createExperimentRequest2);
      Experiment experiment2 = createExperimentResponse2.getExperiment();
      assertEquals(
          "Experiment name not match with expected Experiment name",
          createExperimentRequest2.getName(),
          experiment2.getName());
      LOGGER.info("Create Experiment test successfully executed");
      LOGGER.info("Create Experiment test stop................................");

      LOGGER.info("Create ExperimentRun test start................................");
      CreateExperimentRun createExperimentRunRequest =
          createExperimentRunRequest(project.getId(), experiment.getId());
      CreateExperimentRun.Response createExperimentRunResponse =
          experimentRunServiceStub.createExperimentRun(createExperimentRunRequest);
      ExperimentRun experimentRun = createExperimentRunResponse.getExperimentRun();
      LOGGER.info("ExperimentRun created successfully");
      assertEquals(
          "ExperimentRun name not match with expected ExperimentRun name",
          createExperimentRunRequest.getName(),
          experimentRun.getName());

      CreateExperimentRun createExperimentRunRequest2 =
          createExperimentRunRequest(project.getId(), experiment.getId());
      CreateExperimentRun.Response createExperimentRunResponse2 =
          experimentRunServiceStub.createExperimentRun(createExperimentRunRequest2);
      ExperimentRun experimentRun2 = createExperimentRunResponse2.getExperimentRun();
      LOGGER.info("ExperimentRun created successfully");
      assertEquals(
          "ExperimentRun name not match with expected ExperimentRun name",
          createExperimentRunRequest2.getName(),
          experimentRun2.getName());
      LOGGER.info("Create ExperimentRun test successfully executed");
      LOGGER.info("Create ExperimentRun test stop................................");

      LOGGER.info("Get Project test start................................");
      GetProjects getProjects = GetProjects.newBuilder().build();
      GetProjects.Response getProjectResponse = projectServiceStub.getProjects(getProjects);
      project =
          project
              .toBuilder()
              .setDateUpdated(getProjectResponse.getProjectsList().get(0).getDateUpdated())
              .build();
      assertEquals(
          "Project not match with expected project",
          project,
          getProjectResponse.getProjectsList().get(0));
      LOGGER.info("Get project test successfully executed");
      LOGGER.info("Get project test stop................................");

      LOGGER.info("Get Experiment test start................................");
      GetExperimentById getExperimentRequest =
          GetExperimentById.newBuilder().setId(experiment.getId()).build();
      GetExperimentById.Response getExperimentResponse =
          experimentServiceStub.getExperimentById(getExperimentRequest);
      experiment =
          experiment
              .toBuilder()
              .setDateUpdated(getExperimentResponse.getExperiment().getDateUpdated())
              .build();
      assertEquals(
          "Experiment not match with expected Experiment",
          experiment,
          getExperimentResponse.getExperiment());
      LOGGER.info("Get Experiment test successfully executed");
      LOGGER.info("Get Experiment of project test stop................................");

      LOGGER.info("Get ExperimentRun test start................................");
      GetExperimentRunById request =
          GetExperimentRunById.newBuilder().setId(experimentRun.getId()).build();
      GetExperimentRunById.Response response =
          experimentRunServiceStub.getExperimentRunById(request);
      assertEquals(
          "Experiment not match with expected Experiment",
          experimentRun,
          response.getExperimentRun());
      LOGGER.info("Get ExperimentRun test successfully executed");
      LOGGER.info("Get ExperimentRun test stop................................");

      LOGGER.info("Delete ExperimentRun test start................................");
      DeleteExperimentRun deleteExperimentRunRequest =
          DeleteExperimentRun.newBuilder().setId(experimentRun.getId()).build();
      DeleteExperimentRun.Response deleteExperimentRunResponse =
          experimentRunServiceStub.deleteExperimentRun(deleteExperimentRunRequest);
      assertTrue(deleteExperimentRunResponse.getStatus());
      LOGGER.info("Delete ExperimentRun test successfully executed");
      LOGGER.info("Delete ExperimentRun test stop................................");

      LOGGER.info("Delete Experiment test start................................");
      DeleteExperiment deleteExperimentRequest =
          DeleteExperiment.newBuilder().setId(experiment.getId()).build();
      DeleteExperiment.Response deleteExperimentResponse =
          experimentServiceStub.deleteExperiment(deleteExperimentRequest);
      assertTrue(deleteExperimentResponse.getStatus());
      LOGGER.info("Delete Experiment test successfully executed");
      LOGGER.info("Delete Experiment test stop................................");

      if (isRunningIsolated()) {
        when(uacBlockingMock.getCurrentUser(any())).thenReturn(testUser1);
        mockGetSelfAllowedResources(
            Set.of(project.getId()),
            ModelDBServiceResourceTypes.PROJECT,
            ModelDBServiceActions.DELETE);
      }

      LOGGER.info("Project delete test start................................");
      DeleteProject deleteProjectRequest =
          DeleteProject.newBuilder().setId(project.getId()).build();
      DeleteProject.Response deleteProjectResponse =
          projectServiceStub.deleteProject(deleteProjectRequest);
      assertTrue(deleteProjectResponse.getStatus());
      LOGGER.info("Project delete test successfully executed");
      LOGGER.info("Project delete test stop................................");

    } catch (StatusRuntimeException e) {
      Status status = Status.fromThrowable(e);
      LOGGER.warn("Error Code : " + status.getCode() + " Description : " + status.getDescription());
      fail();
    }

    LOGGER.info("All Entity CRUD Test stop................................");
  }
}
