pipeline {
    agent { label 'docker' }
    stages {
        stage('Build') {
            when {
                branch 'master'
            }
            steps {
                sh "docker build --no-cache --tag dtr.rogfk.no/fint-beta/tech-docs-service:latest ."
            }
        }
        stage('Publish') {
            when {
                branch 'master'
            }
            steps {
                withDockerRegistry([credentialsId: 'dtr-rogfk-no', url: 'https://dtr.rogfk.no']) {
                    sh "docker push dtr.rogfk.no/fint-beta/tech-docs-service:latest"
                }
            }
        }
        stage('Deploy') {
            when {
                branch 'master'
            }
            steps {
                withDockerServer([credentialsId: "ucp-jenkins-bundle", uri: "tcp://ucp.rogfk.no:443"]) {
                    sh "docker service update fint-tech-docs-service_tds --image dtr.rogfk.no/fint-beta/tech-docs-service:latest --detach=false"
                }
            }
        }
    }
}
