pipeline {

    agent any
    environment {
        COMPOSE_PROJECT_NAME = "${env.JOB_NAME}-${env.BUILD_ID}"
        COMPOSE_FILE = "docker-compose.yml"
    }
    stages {
        stage ("Build") {
            steps {
            sh "docker network create -d bridge kong_iafoosball"
            sh "docker-compose build --pull"
            }
        }

        stage ("Deploy") {
            steps {
            sh "docker network create -d bridge kong_iafoosball"
            sh "docker-compose up --force-recreate"
            }
        }
    }
    post {
        always {
            sh "docker-compose down -v --rmi 'all'"
            sh "docker network rm kong_iafoosball"
        
        }
    }
}