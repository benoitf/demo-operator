package appservice

import (
	"context"
	appv1alpha1 "github.com/eivantsov/demo-operator/pkg/apis/app/v1alpha1"
	routev1 "github.com/openshift/api/route/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"time"
)

func (r *ReconcileAppService) createNewPod(instance *appv1alpha1.AppService, pod *corev1.Pod, request reconcile.Request) (err error) {
	reqLogger := log.WithValues("Request.Namespace", request.Namespace, "Request.Name", request.Name)
	if err := controllerutil.SetControllerReference(instance, pod, r.scheme); err != nil {
		return err
	}
	// Check if this Pod already exists
	found := &corev1.Pod{}
	err = r.client.Get(context.TODO(), types.NamespacedName{Name: pod.Name, Namespace: pod.Namespace}, found)
	if err != nil && errors.IsNotFound(err) {
		reqLogger.Info("Creating a new Pod", "Pod.Namespace", pod.Namespace, "Pod.Name", pod.Name)
		err = r.client.Create(context.TODO(), pod)
		if err != nil {
			return err
		}
		// Pod created successfully - don't requeue
		return nil
	}
	return nil
}

func (r *ReconcileAppService) createNewService(instance *appv1alpha1.AppService, service *corev1.Service, request reconcile.Request) (err error) {
	reqLogger := log.WithValues("Request.Namespace", request.Namespace, "Request.Name", request.Name)
	if err := controllerutil.SetControllerReference(instance, service, r.scheme); err != nil {
		reqLogger.Error(err, "An error occurred %s")
		return err
	}
	serviceFound := &corev1.Service{}
	time.Sleep(time.Duration(1) * time.Second)
	err = r.client.Get(context.TODO(), types.NamespacedName{Name: service.Name, Namespace: service.Namespace}, serviceFound)
	if err != nil && errors.IsNotFound(err) {
		reqLogger.Info("Creating a new service", "Service.Namespace", service.Namespace, "Service.Name", service.Name)
		err = r.client.Create(context.TODO(), service)
		if err != nil {
			return err
		}
		return nil
	} else if err != nil {
		return err
	}
	return nil
}

func (r *ReconcileAppService) createNewRoute(instance *appv1alpha1.AppService, route *routev1.Route, request reconcile.Request) (err error)  {
	reqLogger := log.WithValues("Request.Namespace", request.Namespace, "Request.Name", request.Name)
	if err := controllerutil.SetControllerReference(instance, route, r.scheme); err != nil {
		reqLogger.Error(err, "An error occurred %s")
		return err
	}
	routeFound := &routev1.Route{}
	time.Sleep(time.Duration(1) * time.Second)
	err = r.client.Get(context.TODO(), types.NamespacedName{Name: route.Name, Namespace: route.Namespace}, routeFound)
	if err != nil && errors.IsNotFound(err) {
		reqLogger.Info("Creating a new route", "Route.Namespace", route.Namespace, "Route.Name", route.Name)
		err = r.client.Create(context.TODO(), route)
		if err != nil {
			return err
		}
		return nil
	} else if err != nil {
		return err
	}
	return nil
}
