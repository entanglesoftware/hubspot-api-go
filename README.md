# 🚀 HubSpot API Go Client (Golang SDK)

> A powerful, typed, and production-ready **HubSpot API client for Go (Golang)** built by Entangle Software.  
> Integrate with **HubSpot CRM API v3**, Contacts, Deals, Companies, Tickets, and more — cleanly and efficiently.

[![Go Version](https://img.shields.io/badge/go-1.20+-blue.svg)](https://golang.org)
[![License](https://img.shields.io/badge/license-Apache%202.0-green.svg)](LICENSE)
[![HubSpot API](https://img.shields.io/badge/HubSpot-API-orange)](https://developers.hubspot.com/)

---

## ✨ Why This Library?

HubSpot does not provide an official, fully maintained Go SDK covering all APIs with modern typed generation.

This library provides:

- ✅ Typed OpenAPI-generated clients  
- ✅ CRM API v3 support  
- ✅ OAuth 2.0 & Private App token support  
- ✅ Clean client abstraction  
- ✅ Modular architecture  
- ✅ Production-ready structure  
- ✅ Apache 2.0 license  

Perfect for:

- SaaS integrations  
- Backend services  
- Microservices  
- CRM automation tools  
- Enterprise data sync systems  

---

## 📦 Installation

```bash
go get github.com/entanglesoftware/hubspot-api-go@latest
```

---

## 🔐 Authentication

HubSpot supports:

- **Private App Token (Recommended)**
- **OAuth 2.0 Access Token**
- API Key (legacy)

This client supports setting tokens directly.

---

## 🚀 Quick Start

### 1️⃣ Initialize Client

```go
package main

import (
	"fmt"

	"github.com/entanglesoftware/hubspot-api-go/configuration"
	"github.com/entanglesoftware/hubspot-api-go/hubspot"
)

func main() {

	cfg := configuration.Configuration{}

	client := hubspot.NewClient(cfg)

	// Set Access Token (OAuth or Private App)
	client.SetAccessToken("YOUR_ACCESS_TOKEN")

	fmt.Println("HubSpot client initialized successfully")
}
```

---

## 🧩 CRM Usage

The client exposes CRM discovery:

```go
crm := client.Crm()
```

From there, you can access generated endpoint clients inside:

```
codegen/crm/...
```

---

## 📇 Example: Fetch Contacts

```go
package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/entanglesoftware/hubspot-api-go/codegen/crm/objects/contacts"
)

func main() {
	ctx := context.Background()

	client, err := contacts.NewClientWithResponses("https://api.hubapi.com")
	if err != nil {
		panic(err)
	}

	requestEditor := func(ctx context.Context, req *http.Request) error {
		req.Header.Set("Authorization", "Bearer YOUR_ACCESS_TOKEN")
		return nil
	}

	resp, err := client.GetPageWithResponse(ctx, nil, requestEditor)
	if err != nil {
		panic(err)
	}

	fmt.Println("Status:", resp.StatusCode())
}
```

---

## 🏗 Project Structure

```
hubspot/           → Core client wrapper
configuration/     → Config models
discovery/         → CRM discovery
codegen/           → OpenAPI generated clients
oauth/             → OAuth helpers
openapi/           → HubSpot OpenAPI specs
util/              → Shared utilities
```

---

## 🔄 Regenerating API Clients

If OpenAPI specs are updated:

```bash
go generate ./...
```

Generated clients are committed for stability and version control.

---

## ⚙️ Production Best Practices

- Use OAuth or Private App tokens (avoid API keys)
- Inject custom HTTP client with:
  - timeouts
  - retries
  - observability
- Centralize request editor for auth headers
- Log response errors properly

---

## 🌍 Supported APIs

- CRM Contacts
- CRM Deals
- CRM Companies
- CRM Tickets
- Associations
- Marketing APIs (where applicable)

(More endpoints can be generated from OpenAPI specs.)

---

## 🤝 Contributing

Contributions welcome!

1. Fork the repo
2. Create a feature branch
3. Add improvements
4. Submit PR

If you need a new endpoint:
- Open an issue
- Include HubSpot documentation link
- Provide example request/response

---

## 📄 License

Apache License 2.0  
See [LICENSE](LICENSE)

---

## 🔎 SEO Keywords

HubSpot API Go client, HubSpot Go SDK, HubSpot CRM API v3 Golang, HubSpot Contacts API Go, HubSpot OAuth Go, HubSpot integration Golang, HubSpot OpenAPI Go client, HubSpot REST API Go SDK.

---

# ⭐ Star This Repo

If this library helps your integration, please consider giving it a star — it helps the project grow.
