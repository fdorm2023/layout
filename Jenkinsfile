pipeline {
    agent any
    environment {
        IMAGE_NAME = 'fdorm.service.demo'
        VERSION = '0.0.1'
        HTTP_PORT = '8000'
        GRPC_PORT = '9000'
        SERVICE_NAME = 'demo'
	    DOCKERHUB_CREDENTIALS=credentials('docker')
    }
    stages {
        stage('prepare') {
            steps {
                sh "sed -i 's/name_of_service/${SERVICE_NAME}/g' Dockerfile"
                sh "sed -i 's/server_execute_file/${SERVICE_NAME}/g' Dockerfile"
            }
        }
        stage('build') {
            steps {
                sh "docker build -t docker.lvdsoft.com/${IMAGE_NAME}:${VERSION} ."
		        sh 'echo $DOCKERHUB_CREDENTIALS_PSW | docker login docker.lvdsoft.com -u $DOCKERHUB_CREDENTIALS_USR --password-stdin'
                sh "docker push docker.lvdsoft.com/${IMAGE_NAME}:${VERSION}"
            }
        }
        stage('deploy') {
            steps {
                sh "cp app/${SERVICE_NAME}/docker-compose.yaml /home/ubuntu/deploy/${SERVICE_NAME}.yaml"
                    dir('/home/ubuntu/deploy') {
                        sh "docker stack deploy -c ${SERVICE_NAME}.yaml fdorm"
                }
            }
        }
    }
}
