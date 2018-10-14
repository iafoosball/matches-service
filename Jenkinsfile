pipeline {
    agent any
    environment {
        COMPOSE_PROJECT_NAME = "${env.JOB_NAME}-${env.BUILD_ID}"
        COMPOSE_FILE = "docker-compose.yml"
    }
    stages {
        stage ("Prepare environment") {
            steps {
                milestone(ordinal: 1, label: "PREPARE_ENV_MILESTONE")
                sh "rm docker-compose.yml && rm Dockerfile"
                sh "cp ../iaf-configs/matches-service/stag/docker-compose.yml . && cp ../iaf-configs/matches-service/stag/Dockerfile ."
                sh "docker-compose rm -f"
            }
        }
        stage ("Build") {
            steps{
                milestone(ordinal: 2, label: "BUILD_MILESTONE")
                sh "docker-compose build --pull"
                sh "docker cp matches-service:/root/matches.test ."
                sh "docker cp matches-service:/root/maimatches-service ."
            }
        }
        stage ("Test") {
            steps {
                milestone(ordinal: 3, label: "TEST_ENV_MILESTONE")
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