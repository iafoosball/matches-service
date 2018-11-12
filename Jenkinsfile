pipeline {
    agent any

    stages {
        stage ("Prepare stag environment") {
            steps {
                sh "docker stop matches-service-stag &"
                sh "docker stop matches-arangodb-stag &"
                sh "docker rm matches-arangodb-stag &"
                sh "docker rm matches-service-stag &"
            }
        }

        stage ("Build") {
            environment {
                DB_KEY_PROD=credentials('arangoMatchesProd')
            }
            steps{
                sh "printf ${DB_KEY_PROD} >> .env"
                sh "sleep 15s"
                sh "docker-compose build --pull"
                sh blaaa
                sh "sed -i '$ d' .env"
            }
        }

        stage ("Staging") {
            steps {
                sh "docker-compose up -d --force-recreate"
                sh "sleep 30s"
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

                            DB_KEY_PROD=credentials('arangoMatchesProd')

                sh "docker stop matches-service-prod &"
                sh "docker stop matches-arangodb-prod &"
                sh "docker rm matches-arangodb-prod &"
                sh "docker rm matches-service-prod &"
                sh "sleep 15s"
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