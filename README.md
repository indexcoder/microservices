# microservices

# kubectl apply -f .\k8s\broker.yml -- Добавляет новый сервис куберу

# kubectl get pods -- Используется для просмотра списка Pod'ов (контейнеров) в Kubernetes-кластере.

# kubectl get svc -- Используется для просмотра списка сервисов (services) в Kubernetes.

# kubectl delete deployments broker-service mongo rabbitmq  -- Используется для удаления Deployment'ов broker-service, mongo и rabbitmq из текущего namespace.

# kubectl delete svc broker-service mongo rabbitmq -- Используется для удаления сервисов (Service) broker-service, mongo и rabbitmq в текущем namespace.

# kubectl apply -f .\k8s\ -- Используется для применения (деплоя/обновления) всех .yaml-файлов в папке k8s.

# kubectl logs rabbitmq-75d564d49-z7rj8 -- Просмотр логов

# docker-compose -f .\postgres.yml up -d -- Start отдельный контейнер

# minikube dashboard -- Запуск дашбоард
