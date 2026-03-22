package models

import "time"

type User struct {
	ID        string             `json:"id" firestore:"id"`
	Email     string             `json:"email" firestore:"email"`
	CreatedAt time.Time          `json:"createdAt" firestore:"createdAt"`
	UpdatedAt time.Time          `json:"updatedAt" firestore:"updatedAt"`
	Apps      map[string]AppRole `json:"apps" firestore:"apps"`
}

type AppRole struct {
	Role    string    `json:"role" firestore:"role"`
	AddedAt time.Time `json:"addedAt" firestore:"addedAt"`
}

type Role struct {
	ID          string   `json:"id" firestore:"id"`
	Name        string   `json:"name" firestore:"name"`
	Description string   `json:"description" firestore:"description"`
	Permissions []string `json:"permissions" firestore:"permissions"`
}

type App struct {
	ID          string    `json:"id" firestore:"id"`
	Name        string    `json:"name" firestore:"name"`
	Description string    `json:"description" firestore:"description"`
	CreatedAt   time.Time `json:"createdAt" firestore:"createdAt"`
}

type Service struct {
	ID          string            `json:"id" firestore:"id"`
	Name        string            `json:"name" firestore:"name"`
	Description string            `json:"description" firestore:"description"`
	URL         string            `json:"url" firestore:"url"`
	Version     string            `json:"version" firestore:"version"`
	Metadata    map[string]string `json:"metadata" firestore:"metadata"`
	CreatedAt   time.Time         `json:"createdAt" firestore:"createdAt"`
	UpdatedAt   time.Time         `json:"updatedAt" firestore:"updatedAt"`
	IsActive    bool              `json:"isActive" firestore:"isActive"`
}

type ServiceEndpoint struct {
	ID                         string      `json:"id" firestore:"id"`
	ServiceID                  string      `json:"serviceId" firestore:"serviceId"`
	Path                       string      `json:"path" firestore:"path"`
	Method                     string      `json:"method" firestore:"method"` // GET, POST, PUT, DELETE, PATCH
	Summary                    string      `json:"summary" firestore:"summary"`
	InputSchema                interface{} `json:"inputSchema" firestore:"inputSchema"`   // JSON Schema as map[string]interface{}
	OutputSchema               interface{} `json:"outputSchema" firestore:"outputSchema"` // JSON Schema as map[string]interface{}
	AuthRequired               bool        `json:"authRequired" firestore:"authRequired"`
	RateLimitRequestsPerMinute int         `json:"rateLimitRequestsPerMinute" firestore:"rateLimitRequestsPerMinute"`
	RateLimitBurst             int         `json:"rateLimitBurst" firestore:"rateLimitBurst"`
	CreatedAt                  time.Time   `json:"createdAt" firestore:"createdAt"`
	UpdatedAt                  time.Time   `json:"updatedAt" firestore:"updatedAt"`
	IsActive                   bool        `json:"isActive" firestore:"isActive"`
}
