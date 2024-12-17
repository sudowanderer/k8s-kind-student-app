# **Student Management System with Kind and Kubernetes**

A simple **Student Management System** deployed locally on a **Kind**-based Kubernetes cluster. This project demonstrates how to combine a backend (Golang + MongoDB), a frontend (Nginx), and Kubernetes manifests to set up a microservice application for local development and testing.

------

## **Project Features**

- 🛠 **Backend**: Developed in **Golang** with RESTful APIs.
- 🌐 **Frontend**: Static files served via **Nginx**.
- 🐳 **Containerized**: Dockerized application services.
- 🚀 **Kubernetes Deployment**: Using **Kind** to run a local K8s cluster.
- 📦 **Database**: MongoDB as the backend database.

Before running the project, ensure you have the following tools installed:

- **Docker**: To build and run the application containers.
- **Kind**: To set up a local Kubernetes cluster.
- **kubectl**: To interact with the Kubernetes cluster.

------

## **Setup Instructions**

### **1. Set Up Kind Cluster**

Create a local Kubernetes cluster using Kind:

```bash
kind create cluster --name student-cluster --config kind-cluster-config.yaml
```

### **2. Build Docker Images**

Build the backend and frontend Docker images:

```bash
# Build backend
docker build -t student-backend:latest ./backend
```

### **3. Load Docker Images into Kind Cluster**

将 Docker 镜像加载到 Kind 集群中：

```bash
kind load docker-image student-backend:latest -n student-cluster
```

### **4. Install Ingress-NGINX**

在 Kind 集群中安装 Ingress-NGINX 控制器：

```bash
kubectl apply -f https://kind.sigs.k8s.io/examples/ingress/deploy-ingress-nginx.yaml
```

验证 Ingress 控制器是否正常运行：

```bash
kubectl get pods -n ingress-nginx
```

------

### **5. Deploy Application to Kubernetes**

Apply the Kubernetes manifests to deploy the services:

```bash
kubectl apply -f k8s
```

------

### **6. Verify Deployments**

检查 Pod 和服务的状态：

```bash
kubectl get pods
kubectl get svc
```

------

## **Access the Application**

一旦所有服务运行正常，通过浏览器访问应用：

- **Frontend**: `http://localhost`

- Backend API

  : 测试 API 路径，例如：

  ```bash
  curl http://localhost/api/v1/students
  ```

------

## **Endpoints**

| **Method** | **Endpoint**       | **Description**   |
| ---------- | ------------------ | ----------------- |
| GET        | `/api/v1/students` | Get all students  |
| POST       | `/api/v1/students` | Add a new student |

------

## **Clean Up**

删除 Kind 集群并释放资源：

```bash
kind delete cluster --name student-cluster
```
