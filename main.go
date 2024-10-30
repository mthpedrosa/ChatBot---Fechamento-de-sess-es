package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/robfig/cron/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Função que fecha sessões abertas há mais de 24 horas
func closeOldSessions(ctx context.Context, db *mongo.Database) {
	collection := db.Collection("sessions")
	filter := bson.M{
		"created_at": bson.M{"$lte": time.Now().Add(-24 * time.Hour)},
		"status":     "in_progress", // status para sessões ainda abertas
	}
	update := bson.M{"$set": bson.M{"status": "closed", "finished_at": time.Now()}}

	result, err := collection.UpdateMany(ctx, filter, update)
	if err != nil {
		fmt.Printf("Error closing old sessions: %v\n", err)
		return
	}

	fmt.Printf("Closed %d sessions older than 24 hours\n", result.ModifiedCount)
}

func main() {
	// Inicia o cron scheduler
	c := cron.New()

	// Contexto para operações MongoDB
	ctx := context.Background()

	// Configuração do cliente MongoDB
	mongoURI := os.Getenv("MONGO_URI")
	clientOptions := options.Client().ApplyURI(mongoURI)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		fmt.Printf("Failed to connect to MongoDB: %v\n", err)
		return
	}
	defer client.Disconnect(ctx)

	db := client.Database("autflow")

	// Executa `closeOldSessions` uma vez agora
	fmt.Println("Running closeOldSessions now...")
	closeOldSessions(ctx, db)

	// Agendamento da função para rodar todos os dias à meia-noite
	c.AddFunc("@daily", func() {
		closeOldSessions(ctx, db)
	})

	// Inicia o cron scheduler
	c.Start()

	// Captura sinais para finalizar o cron e desconectar do MongoDB
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	fmt.Println("Cron job started. Press Ctrl+C to exit.")

	<-sig // Espera até o usuário interromper a execução
	fmt.Println("Shutting down...")
	c.Stop()
}
