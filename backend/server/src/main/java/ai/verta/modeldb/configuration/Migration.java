package ai.verta.modeldb.configuration;

import ai.verta.modeldb.common.MssqlMigrationUtil;
import ai.verta.modeldb.common.futures.FutureJdbi;
import ai.verta.modeldb.config.MDBConfig;
import ai.verta.modeldb.utils.ModelDBHibernateUtil;
import java.util.Optional;
import org.apache.logging.log4j.LogManager;
import org.apache.logging.log4j.Logger;

public class Migration {
  private final Logger LOGGER = LogManager.getLogger(Migration.class);

  public Migration(MDBConfig mdbConfig, FutureJdbi futureJdbi) throws Exception {
    migrate(mdbConfig, futureJdbi);
  }

  public void migrate(MDBConfig config, FutureJdbi futureJdbi) throws Exception {
    var databaseConfig = config.getDatabase();
    var modelDBHibernateUtil = ModelDBHibernateUtil.getInstance();
    modelDBHibernateUtil.initializedConfigAndDatabase(config, databaseConfig);

    LOGGER.info("Starting Migrations");
    modelDBHibernateUtil.runMigrations(databaseConfig, "migration", Optional.of(1));
    LOGGER.info("Migrations Complete");

    // Initialized session factory before code migration and after liquibase migration
    modelDBHibernateUtil.createOrGetSessionFactory(databaseConfig);

    LOGGER.info("Code migration starting");
    modelDBHibernateUtil.runMigration(databaseConfig, config.getMigrations());
    LOGGER.info("Code migration done");

    if (databaseConfig.getRdbConfiguration().isMssql()) {
      MssqlMigrationUtil.migrateToUTF16ForMssql(futureJdbi);
    }
  }
}
