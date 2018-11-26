pipeline {
    environment {
        PW_STAG=credentials('arangoMatchesStag')
        PW_PROD=credentials('arangoMatchesProd')
    }
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
            steps{
                sh "printf arangoPasswordStag=${PW_STAG} >> .env"
                sh "sleep 3s"
                sh "docker-compose -f docker-compose.yml -f docker-compose.stag.yml build"
            }
        }

        stage ("Staging") {
            steps {
                sh "docker-compose -f docker-compose.yml -f docker-compose.stag.yml up -d --force-recreate"
                sh "sleep 30s"
            }
        }

        stage ("Test") {
            steps {
                sh "docker exec matches-service-stag /matches.test --dbhost=matches-arangodb-stag --dbport=8529"
            }
        }

        stage ("Prepare prod environment") {
            steps {
                sh "printf arangoPasswordStag=${PW_STAG} >> .env"
                sh "sleep 3s"
                sh "docker-compose -f docker-compose.yml -f docker-compose.prod.yml build"
            }
        }

        stage ("Production") {





           environment {
                           DB_PW_Stag=credentials('arangoMatchesStag')
                       }
                       steps{
                           sh "sed -i '\$ d' .env"
                           sh "printf ${DB_PW_STAG} >> .env"
                           sh "docker-compose -f docker-compose.yml -f docker-compose.stag.yml build"
                           sh "sed -i '\$ d' .env"





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