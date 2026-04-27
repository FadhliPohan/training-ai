# 📋 Simplified Sprint Implementation Plan

> **Project:** InsightFlow - Self-Service AI Dashboard for Clothing Sales  
> **Developer:** Sandi (Software Engineer with 12 years experience)  
> **Start Date:** April 27, 2026  
> **Target MVP:** July 17, 2026 (12 weeks)

---

## 🎯 Project Goals

1. **Reduce IT Dependency** - >70% reduction in ad-hoc report requests
2. **Faster Anomaly Response** - Days → Minutes
3. **24/7 Customer Service** - AI Chat Assistant availability
4. **Mobile-First Monitoring** - Telegram daily summaries

---

## 🚀 Overall Timeline (12 Weeks)

| Sprint | Dates | Focus | Goal |
|---|---|---|---|
| **Sprint 1** | Apr 27 - May 8 | Backend Auth + Master Data | Authentication system working |
| **Sprint 2** | May 11 - May 22 | Backend Transactions | Order processing complete |
| **Sprint 3** | May 25 - Jun 5 | Backend Reports + AI | Dashboard analytics working |
| **Sprint 4** | Jun 8 - Jun 19 | Frontend Foundation | UI basics functional |
| **Sprint 5** | Jun 22 - Jul 3 | Frontend Dashboard | Full dashboard UI complete |
| **Sprint 6** | Jul 6 - Jul 17 | Security + Testing | Production ready |

---

## 📌 Current Status: Sprint 1 (In Progress)

### ✅ Completed (Week 1 - Apr 27-28)
- [x] Project setup and environment configuration
- [x] Authentication system (`POST /auth/login`)
- [x] JWT middleware implementation
- [x] Swagger API documentation
- [x] DTO creation for all modules
- [x] Database migration system (automatic)

### 🔄 In Progress (Week 1 - Apr 29-30)
- [ ] `POST /auth/logout` endpoint
- [ ] `GET /auth/me` endpoint
- [ ] Role-based access control middleware
- [ ] Product CRUD operations (GET, POST)
- [ ] Customer CRUD operations (GET, POST)

### 🔜 Coming Up (Week 2 - May 1-8)
- [ ] Complete Product CRUD (PUT, PATCH)
- [ ] Complete Customer CRUD (PUT)
- [ ] User management endpoints
- [ ] API documentation completion
- [ ] Sprint 1 Review & Demo

---

## 🛠️ Daily Development Workflow

### 1. Start Your Day
```bash
# 1. Start PostgreSQL (if not already running)
cd /home/sandi/PUSRI/training-ai
docker compose up -d postgres

# 2. Navigate to backend
cd /home/sandi/PUSRI/training-ai/be-penjualan

# 3. Run the application
make run
# or: go run cmd/main.go
```

### 2. Test Your Work
```bash
# Test health endpoint
curl http://localhost:8080/health

# Test login (existing)
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email":"admin@insightflow.id","password":"Admin@12345"}'

# Check Swagger UI
# Open: http://localhost:8080/swagger/index.html
```

### 3. Development Commands
```bash
# Run with specific environment
make dev        # Development mode
make staging    # Staging mode
make prod       # Production mode

# Run tests
make test

# Create new migration
make create-migration
```

---

## 📁 Project Structure You're Working With

```
/home/sandi/PUSRI/training-ai/be-penjualan/
├── cmd/main.go                 # Application entry point
├── internal/
│   ├── handler/               # API endpoints
│   │   └── auth/             # Authentication endpoints
│   │       ├── auth.go       # Handler setup
│   │       └── login.go      # Login implementation
│   ├── dto/                  # Data Transfer Objects
│   ├── domain/               # Domain models
│   ├── database/             # Database connection
│   ├── middleware/           # Auth middleware
│   └── router/               # Route configuration
├── db/migrations/            # Database migration files
├── docs/                     # Swagger documentation
├── .env                      # Environment configuration
└── Makefile                  # Development commands
```

---

## 🎯 Sprint 1 Daily Targets

### Day 1-2 (Apr 27-28) - COMPLETED ✓
- Authentication system
- JWT implementation
- Documentation setup

### Day 3-4 (Apr 29-30) - CURRENTLY WORKING
- [ ] Task 1: `POST /auth/logout` endpoint
- [ ] Task 2: `GET /auth/me` endpoint  
- [ ] Task 3: RoleGuard middleware
- [ ] Task 4: Product GET endpoint
- [ ] Task 5: Product POST endpoint

### Day 5-6 (May 1-2) - NEXT TARGET
- [ ] Task 6: Product PUT/PATCH endpoints
- [ ] Task 7: Customer CRUD endpoints
- [ ] Task 8: User management endpoints

### Day 7-8 (May 5-6) - FINALIZING
- [ ] Task 9: API documentation completion
- [ ] Task 10: Testing and bug fixes
- [ ] Task 11: Code review preparation

### Day 9-10 (May 7-8) - WRAP UP
- [ ] Sprint 1 Review preparation
- [ ] Demo ready
- [ ] Sprint 2 Planning

---

## 📊 Progress Tracking

### Overall Progress
- **Project Completion:** ~7%
- **Sprint 1 Progress:** 40% (8/20 tasks)
- **Next Major Milestone:** May 8 - Sprint 1 Complete

### Daily Checklist
- [ ] Start PostgreSQL container
- [ ] Run application
- [ ] Complete assigned tasks
- [ ] Test endpoints
- [ ] Update documentation
- [ ] Commit working code

---

## 🚨 Quick Troubleshooting

### Common Issues and Solutions

1. **Database Connection Failed**
   ```bash
   # Make sure PostgreSQL is running
   docker compose up -d postgres
   docker compose ps  # Should show postgres running
   ```

2. **JWT Secret Error**
   ```bash
   # Make sure .env file has proper JWT_SECRET (32+ characters)
   # Or run with: unset JWT_SECRET && go run cmd/main.go
   ```

3. **Port Already in Use**
   ```bash
   # Kill existing process
   lsof -i :8080
   kill -9 <PID>
   ```

4. **Migration Issues**
   ```bash
   # Reset database (development only)
   docker compose down
   docker compose up -d postgres
   ```

---

## 📞 When You Need Help

### For Code Issues
1. Check existing implementation in `internal/handler/auth/`
2. Look at DTO structures in `internal/dto/`
3. Refer to middleware in `internal/middleware/`

### For Environment Issues
1. Verify `.env` file configuration
2. Check PostgreSQL connection with:
   ```bash
   PGPASSWORD=insightflow123 psql -h localhost -p 5433 -U insightflow -d insightflow_db
   ```

### For API Documentation
1. Access Swagger at: `http://localhost:8080/swagger/index.html`
2. Check existing annotations in handler files

---

## 🎉 Success Criteria for Sprint 1

By May 8, 2026, you should have:

1. ✅ **Authentication System**
   - Login endpoint working
   - Logout endpoint working
   - User profile endpoint working

2. ✅ **Authorization System**
   - Role-based access control
   - Read-only restrictions for viewers

3. ✅ **Master Data Management**
   - Product CRUD operations
   - Customer CRUD operations
   - User management operations

4. ✅ **Documentation**
   - Swagger UI working
   - All endpoints documented
   - Examples provided

5. ✅ **Infrastructure**
   - Automatic database migrations
   - Development workflow established
   - Testing capability

---

## 📅 Next Sprint Preview

**Sprint 2 (May 11-22): Transaction Processing**
- Order creation and management
- Payment processing
- Shipment tracking
- Data validation and transactions

This simplified plan focuses only on what you need to know RIGHT NOW for Sprint 1. Everything else has been removed to avoid confusion.