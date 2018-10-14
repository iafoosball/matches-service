pipeline {
    agent any
    environment {
        COMPOSE_FILE = "docker-compose.yml"
    }

    stages {
        stage ("Prepare stag environment") {
            steps {
                sh "rm docker-compose.yml && rm Dockerfile"
                sh "cp ../iaf-configs/matches-service/stag/docker-compose.yml . && cp ../iaf-configs/matches-service/stag/Dockerfile ."

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
                sh "docker stop matches-service-stag &"
                sh "docker stop matches-arangodb-stag &"
                sh "docker-compose up --force-recreate -d"
                sh "sleep 30s"
                sh "docker cp matches-service:/root/matches.test ."
                sh "./matches.test"
                sh "docker-compose down"
            }
        }
        stage ("Staging") {
                    steps {
                        sh "docker-compose up -d"
                    }
                }
        stage ("Prepare prod environment") {
                    steps {
                        sh "rm docker-compose.yml && rm Dockerfile"
                        sh "cp ../iaf-configs/matches-service/prod/docker-compose.yml . && cp ../iaf-configs/matches-service/prod/Dockerfile ."
                    }
                }
        stage ("Production") {
            steps {
                sh "docker stop matches-service-prod &"
                sh "docker stop matches-arangodb-prod &"
                sh "docker-compose up --force-recreate"
            }
        }
    }
    post {
        always {
            sh "docker-compose down -v --rmi 'all'"
        }
    }
}