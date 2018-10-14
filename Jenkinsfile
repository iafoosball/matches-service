pipeline {
    agent any
    environment {
        COMPOSE_FILE = "docker-compose.yml"
    }

    stages {
        stage ("Prepare environment") {
            steps {
                sh "rm docker-compose.yml && rm Dockerfile"
                sh "cp ../iaf-configs/matches-service/stag/docker-compose.yml . && cp ../iaf-configs/matches-service/stag/Dockerfile ."
                sh "docker stop matches-service  && docker stop matches-arangodb"
                sh "docker-compose rm -f"
            }
        }
        stage ("Build") {
            steps{
                sh "docker-compose build --pull"

            }
        }
        stage ("Test") {
            steps {
                sh "docker-compose up --force-recreate -d"
                sh "sleep 30s"
                sh "docker cp matches-service:/root/matches.test ."
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
            sh "docker-compose down -v --rmi 'all'"
        }
    }
}