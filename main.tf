terraform {
  required_providers {
    kubernetes = {
      source  = "hashicorp/kubernetes"
      version = "~> 2.27.0"
    }
  }
}

resource "kubernetes_deployment" "cats" {
  metadata {
    name      = "catregistry-deployment"
    namespace = var.namespace
  }
  spec {
    replicas = 3
    selector {
      match_labels = {
        app  = "catregistry"
        tier = "backend"
        role = "api"
      }
    }
    template {
      metadata {
        namespace = var.namespace
        labels = {
          app  = "catregistry"
          tier = "backend"
          role = "api"
        }
      }
      spec {
        container {
          name  = "cat-container"
          image = "oksuriini/catregistry:0.1"
          env {
            name  = "MONGODB_URI"
            value = "mongodb-service"
          }
          port {
            container_port = 8080
          }
        }
      }
    }
  }
}

resource "kubernetes_stateful_set" "mongodb" {
  metadata {
    name      = "mongodb-deployment"
    namespace = var.namespace
  }
  spec {
    service_name = "mongodb-serv"
    replicas     = 1
    selector {
      match_labels = {
        app  = "catregistry"
        tier = "backend"
        role = "database"
      }
    }
    template {
      metadata {
        namespace = var.namespace
        labels = {
          app  = "catregistry"
          tier = "backend"
          role = "database"
        }
      }
      spec {
        container {
          name  = "mongodb-container"
          image = "mongo"
          port {
            container_port = 27017
            name           = "mongodb"
          }
        }
      }
    }
  }
}

resource "kubernetes_service" "backend" {
  metadata {
    name      = "mongodb-service"
    namespace = var.namespace
  }
  spec {
    selector = {
      app  = "catregistry"
      tier = "backend"
      role = "database"
    }
    port {
      port        = 27017
      target_port = 27017
    }
  }
}

resource "kubernetes_service" "goback" {
  metadata {
    name      = "gobackend-service"
    namespace = var.namespace
  }
  spec {
    type = "NodePort"
    port {
      port        = 8080
      target_port = 8080
    }
    selector = {
      app  = "catregistry"
      tier = "backend"
      role = "api"
    }
  }
}
