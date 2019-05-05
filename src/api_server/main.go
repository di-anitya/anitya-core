package main

import (
	"fmt"
	"net/http"
	"os"

	ctrv1 "./controllers/v1"
	handler "./handler"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()
	router.Use(handler.JwtAuthentication) //attach JWT auth middleware

	v1 := router.PathPrefix("/api/v1").Subrouter()

	v1.HandleFunc("/auth/tokens", ctrv1.PublishToken).Methods("POST")

	v1.HandleFunc("/project", ctrv1.CreateProject).Methods("POST")
	v1.HandleFunc("/project", ctrv1.ListProject).Methods("GET")
	v1.HandleFunc("/project/{project_id}", ctrv1.ShowProject).Methods("GET")
	v1.HandleFunc("/project/{project_id}", ctrv1.ModifyAndShowProject).Methods("PATCH")
	v1.HandleFunc("/project/{project_id}", ctrv1.DeleteProject).Methods("DELETE")

	v1.HandleFunc("/role", ctrv1.CreateRole).Methods("POST")
	v1.HandleFunc("/role", ctrv1.ListRole).Methods("GET")
	v1.HandleFunc("/role/{role_id}", ctrv1.ShowRole).Methods("GET")
	v1.HandleFunc("/role/{role_id}", ctrv1.ModifyAndShowRole).Methods("PATCH")
	v1.HandleFunc("/role/{role_id}", ctrv1.DeleteRole).Methods("DELETE")

	v1.HandleFunc("/user", ctrv1.CreateUser).Methods("POST")
	v1.HandleFunc("/user", ctrv1.ListUser).Methods("GET")
	v1.HandleFunc("/user/{user_id}", ctrv1.ShowUser).Methods("GET")
	v1.HandleFunc("/user/{user_id}", ctrv1.ModifyAndShowUser).Methods("PATCH")
	v1.HandleFunc("/user/{user_id}", ctrv1.DeleteUser).Methods("DELETE")

	v1.HandleFunc("/dashboard", ctrv1.CreateDashboardConfig).Methods("POST")
	v1.HandleFunc("/dashboard/{dashboard_id}", ctrv1.ShowDashboardConfig).Methods("GET")
	v1.HandleFunc("/dashboard/{dashboard_id}", ctrv1.ModifyAndShowDashboardConfig).Methods("PATCH")
	v1.HandleFunc("/dashboard/{dashboard_id}", ctrv1.DeleteDashboardConfig).Methods("DELETE")

	v1.HandleFunc("/probepoint", ctrv1.CreateProbePoint).Methods("POST")
	v1.HandleFunc("/probepoint", ctrv1.ListProbePoint).Methods("GET")
	v1.HandleFunc("/probepoint/{probepoint_id}", ctrv1.ShowProbePoint).Methods("GET")
	v1.HandleFunc("/probepoint/{probepoint_id}", ctrv1.ModifyAndShowProbePoint).Methods("PATCH")
	v1.HandleFunc("/probepoint/{probepoint_id}", ctrv1.DeleteProbePoint).Methods("DELETE")

	v1.HandleFunc("/monitoring/http", ctrv1.CreateHTTPMonitoringConfig).Methods("POST")
	v1.HandleFunc("/monitoring/http", ctrv1.ListHTTPMonitoringConfig).Methods("GET")
	v1.HandleFunc("/monitoring/http/{http_id}", ctrv1.ShowHTTPMonitoringConfig).Methods("GET")
	v1.HandleFunc("/monitoring/http/{http_id}", ctrv1.ModifyAndShowHTTPMonitoringConfig).Methods("PATCH")
	v1.HandleFunc("/monitoring/http/{http_id}", ctrv1.DeleteHTTPMonitoringConfig).Methods("DELETE")

	v1.HandleFunc("/monitoring/dns", ctrv1.CreateDNSMonitoringConfig).Methods("POST")
	v1.HandleFunc("/monitoring/dns", ctrv1.ListDNSMonitoringConfig).Methods("GET")
	v1.HandleFunc("/monitoring/dns/{dns_id}", ctrv1.ShowDNSMonitoringConfig).Methods("GET")
	v1.HandleFunc("/monitoring/dns/{dns_id}", ctrv1.ModifyAndShowDNSMonitoringConfig).Methods("PATCH")
	v1.HandleFunc("/monitoring/dns/{dns_id}", ctrv1.DeleteDNSMonitoringConfig).Methods("DELETE")

	v1.HandleFunc("/notification/history", ctrv1.CreateNotificationHistory).Methods("POST")
	v1.HandleFunc("/notification/history", ctrv1.ListNotificationHistory).Methods("GET")
	v1.HandleFunc("/notification/history/{notification_id}", ctrv1.ShowNotificationHistory).Methods("GET")
	v1.HandleFunc("/notification/history/{notification_id}", ctrv1.ModifyAndShowNotificationHistory).Methods("PATCH")
	v1.HandleFunc("/notification/history/{notification_id}", ctrv1.DeleteNotificationHistory).Methods("DELETE")

	v1.HandleFunc("/notification/slack", ctrv1.CreateNotificationSlack).Methods("POST")
	v1.HandleFunc("/notification/slack", ctrv1.ListNotificationSlack).Methods("GET")
	v1.HandleFunc("/notification/slack/{notification_id}", ctrv1.ShowNotificationSlack).Methods("GET")
	v1.HandleFunc("/notification/slack/{notification_id}", ctrv1.ModifyAndShowNotificationSlack).Methods("PATCH")
	v1.HandleFunc("/notification/slack/{notification_id}", ctrv1.DeleteNotificationSlack).Methods("DELETE")

	port := os.Getenv("PORT") //Get port from .env file, we did not specify any port so this should return an empty string when tested locally
	if port == "" {
		port = "8080" //localhost
	}

	fmt.Println("api-server is running.. (0.0.0.0:" + port + ")")

	err := http.ListenAndServe(":"+port, router) //Launch the app, visit localhost:8000/api
	if err != nil {
		fmt.Print(err)
	}
}
