pipeline {
    @NonCPS
    void cancelPreviousRunningBuilds(int maxBuildsToSearch = 20) {
        RunWrapper b = currentBuild
        for (int i=0; i<maxBuildsToSearch; i++) {
            b = b.getPreviousBuild();
            if (b == null) break;
            Run<?,?> rawBuild = b.rawBuild
            if (rawBuild.isBuilding()) {
                rawBuild.doStop()
            }
        }
    }
    agent any
    environment {
        COMPOSE_PROJECT_NAME = "${env.JOB_NAME}-${env.BUILD_ID}"
        COMPOSE_FILE = "docker-compose.yml"
    }
    options {
            disableConcurrentBuilds()
    }
    stages {
        stage ("Prepare environment") {
            steps {
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
               cancelPreviousBuilds()
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