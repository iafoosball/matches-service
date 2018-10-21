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
            }
        }
        stage ("Build") {
            steps{
                sh "docker stop matches-service-stag &"
                sh "docker stop matches-arangodb-stag &"
                sh "sleep 15s"
                sh "docker-compose build --pull"

            }
        }
        stage ("Staging") {
                    steps {
                        sh "docker-compose up -d --force-recreate"
                        sh "sleep 100s"
                    }
                }
        stage ("Test") {
            steps {
                sh "docker cp matches-service-stag:/root/matches.test ."
                sh "./matches.test"
            }
        }
        stage ("Prepare prod environment") {
                    steps {
                        sh "rm docker-compose.yml && rm Dockerfile"
                        sh "cp ../iaf-configs/matches-service/prod/docker-compose.yml . && cp ../iaf-configs/matches-service/prod/Dockerfile ."
                        sh "cp -rf matches.yml /var/lib/iafoosball/swagger-ui/ &"
                    }
                }
        stage ("Production") {
            steps {
                sh "docker stop matches-service-prod &"
                sh "docker stop matches-arangodb-prod &"
                sh "docker-compose up --force-recreate --build"
            }
        }
    }
    post {
        always {
            sh "docker-compose down -v --rmi 'all'"
        }
    }
}