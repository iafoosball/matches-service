node {

    stage('Initialize'){
        def dockerHome = tool 'docker'
        env.PATH = "${dockerHome}/bin:${env.PATH}"
    }


    stage('Build'){
        sh "docker-compose up"
    }
    }