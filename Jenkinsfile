def cancelPreviousBuilds() {
    def jobName = env.JOB_NAME
    def buildNumber = env.BUILD_NUMBER.toInteger()
    /* Get job name */
    def currentJob = Jenkins.instance.getItemByFullName(jobName)

    /* Iterating over the builds for specific job */
    for (def build : currentJob.builds) {
        /* If there is a build that is currently running and it's not current build */
        if (build.isBuilding() && build.number.toInteger() != buildsNumber) {
            /* Than stopping it */
            build.doKill()
        }
    }
}
pipeline {
    agent any
    environment {
        COMPOSE_PROJECT_NAME = "${env.JOB_NAME}-${env.BUILD_ID}"
        COMPOSE_FILE = "docker-compose.yml"
    }
    stages {
        stage ("Prepare environment") {
            steps {
                def jobname = env.JOB_NAME
                echo jobname
                def job = Jenkins.instance.getItemByFullName(jobname)
                for (def build : currentJob.builds) {
                    /* If there is a build that is currently running and it's not current build */
                    if (build.isBuilding() && build.number.toInteger() != buildsNumber) {
                           sh "curl -X GET http://localhost:8030/matches $build.number.toInteger()"
                    }
                }

                sh "rm docker-compose.yml && rm Dockerfile"
                sh "cp ../iaf-configs/matches-service/stag/docker-compose.yml . && cp ../iaf-configs/matches-service/stag/Dockerfile ."
                sh "docker-compose rm -f"
            }
        }
        stage ("Build") {
            steps{
                sh "docker-compose build --pull"
                sh "docker cp matches-service:/root/matches.test ."
                sh "docker cp matches-service:/root/maimatches-service ."
            }
        }
        stage ("Test") {
            steps {
                sh "docker-compose up --force-recreate -d"
                sh "sleep 30s"
                sh "./matches.test"
                sh "docker-compose down"
            }
        }
        stage ("Production") {
            steps {
                sh "docker-compose up"
            }
        }
    }
    post {
        always {
            sh "docker-compose down -v"
        }
    }
}