node {

    stage('Initialize'){
        def dockerHome = tool 'myDocker'
        env.PATH = "${dockerHome}/bin:${env.PATH}"
    }


    stage('Build'){
        sh "docker-compose up"
    }
    }