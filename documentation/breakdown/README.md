# InsightFlow — Work Breakdown Index

> **Proyek:** InsightFlow Self-Service AI Dashboard Penjualan Pakaian  
> **Versi Dokumen:** v2.0  
> **Terakhir Diperbarui:** 24 April 2026

---

## 📚 Dokumentasi Lengkap

### 🎯 Sprint Planning & Execution

| Dokumen | Tujuan | Kapan Digunakan |
|---|---|---|
| **[SIMPLIFIED_SPRINT_PLAN.md](../SIMPLIFIED_SPRINT_PLAN.md)** | **Panduan harian Sprint 1 yang disederhanakan** | **Untuk Backend Developer (Anda!)** |

---

## 📋 Work Breakdown Structure

### Area Breakdown (Technical)

| Area | File | Deskripsi | Status |
|---|---|---|---|
| **Area 1** | [area-1-api-documentation.md](./area-1-api-documentation.md) | Kontrak API, format standar, 21 endpoint | 🟡 In Progress |
| **Area 2** | [area-2-product-requirements.md](./area-2-product-requirements.md) | Klarifikasi PRD, 8 laporan dashboard | ✅ Complete |
| **Area 3** | [area-3-backend-architecture.md](./area-3-backend-architecture.md) | Go Fiber, DB migration, Auth, CRUD | 🟡 In Progress |
| **Area 4** | [area-4-frontend-uiux.md](./area-4-frontend-uiux.md) | Next.js 14, Design System, 20+ halaman | 🔴 Not Started |
| **Area 5** | [area-5-security.md](./area-5-security.md) | Auth, Authorization, Security hardening | 🔴 Not Started |
| **Area 6** | [area-6-n8n-ai-workflows.md](./area-6-n8n-ai-workflows.md) | 5 workflow n8n + AI integration | 🔴 Not Started |

---

## 🗺️ Quick Navigation

```
documentation/
├── README.md                      ← You are here
├── Self_Service_AI_Dashboard_PRD.md
├── work_breakdown.md
├── task_plan.md
├── implementation_sprint_plan.md
│
├── 🆕 SPRINT_IMPLEMENTATION_PLAN.md    ← Master sprint plan (6 sprints)
├── 🆕 SPRINT_TRACKER.md                ← Daily progress tracker
├── 🆕 SPRINT_1_KICKOFF.md              ← Sprint 1 execution guide
│
└── breakdown/
    ├── area-1-api-documentation.md
    ├── area-2-product-requirements.md
    ├── area-3-backend-architecture.md
    ├── area-4-frontend-uiux.md
    ├── area-5-security.md
    └── area-6-n8n-ai-workflows.md
```

---

## 📊 Project Status Overview

### Current Sprint: **Sprint 1** (27 Apr - 08 Mei 2026)

**Theme:** Backend Auth + Master Data  
**Progress:** ~5% complete  
**Focus:** Authentication system, Role-based access, Produk & Customer CRUD

| Phase | Completion | Status |
|---|---|---|
| Sprint 0 (Foundation) | 100% | ✅ Complete |
| Sprint 1 (Auth + Master) | 5% | 🟡 In Progress |
| Sprint 2 (Transactions) | 0% | 🔴 Not Started |
| Sprint 3 (Reports + n8n) | 0% | 🔴 Not Started |
| Sprint 4 (Frontend) | 0% | 🔴 Not Started |
| Sprint 5 (Dashboard) | 0% | 🔴 Not Started |
| Sprint 6 (Security + UAT) | 0% | 🔴 Not Started |

**Overall Project Completion:** ~5%

---

## 🎯 Key Milestones

| Date | Milestone | Owner |
|---|---|---|
| 26 Apr 2026 | ✅ Sprint 0 Complete (Foundation) | Backend Team |
| 08 Mei 2026 | Sprint 1 Review + Demo | All Team |
| 22 Mei 2026 | Sprint 2 Complete (Transactions) | Backend Team |
| 05 Jun 2026 | Sprint 3 Complete (Reports + n8n) | Backend + AI |
| 19 Jun 2026 | Sprint 4 Complete (Frontend Foundation) | Frontend Team |
| 03 Jul 2026 | Sprint 5 Complete (Dashboard AI) | All Team |
| 17 Jul 2026 | **Go-Live Decision** | Stakeholders |

---

## 🔗 External Resources

| Resource | Link |
|---|---|
| Backend Repository | `/home/sandi/PUSRI/training-ai/be-penjualan/` |
| Frontend Repository | `/home/sandi/PUSRI/training-ai/fe-penjualan/` |
| AGENTS.md (Project Rules) | `/home/sandi/PUSRI/training-ai/be-penjualan/.agents/AGENTS.md` |

---

## 📞 Team Contacts

| Role | Responsibilities | Sprint Focus |
|---|---|---|
| Backend Lead | API design, architecture, code review | S1-S3, S6 |
| Backend Developer | Implementation, testing | S1-S3, S6 |
| Frontend Lead | UI architecture, components | S4-S6 |
| Frontend Developer | Implementation | S4-S6 |
| DevOps | Infrastructure, CI/CD, deployment | S0, S3, S6 |
| n8n/AI Specialist | Workflow creation, LLM integration | S3, S5 |
| Product Manager | Stakeholder management, prioritization | All sprints |

---

## 🚀 Getting Started

### For New Team Members

1. **Read the PRD** → [Self_Service_AI_Dashboard_PRD.md](../Self_Service_AI_Dashboard_PRD.md)
2. **Review Sprint Plan** → [SPRINT_IMPLEMENTATION_PLAN.md](../SPRINT_IMPLEMENTATION_PLAN.md)
3. **Check Daily Tracker** → [SPRINT_TRACKER.md](../SPRINT_TRACKER.md)
4. **Follow Area Breakdown** → Start with your area (Backend → Area 3, Frontend → Area 4)

### For Backend Developers

Start here: [SPRINT_1_KICKOFF.md](../SPRINT_1_KICKOFF.md)

### For Frontend Developers

Start here: Area 4 breakdown (will be activated in Sprint 4)

### For Stakeholders

Review: [SPRINT_IMPLEMENTATION_PLAN.md](../SPRINT_IMPLEMENTATION_PLAN.md) (Section 1 & 2)  
Check progress: [SPRINT_TRACKER.md](../SPRINT_TRACKER.md) (Updated twice weekly)

---

## 📝 Document History

| Version | Date | Changes | Author |
|---|---|---|---|
| v3.0 | 29 Apr 2026 | Simplified for developer focus | Sandi |
| v2.0 | 24 Apr 2026 | Added sprint planning documents | Engineering Team |
| v1.0 | Apr 2026 | Initial work breakdown structure | Engineering Team |

---

## ✅ Next Actions

**Today (29 Apr 2026):**
- [ ] Complete `POST /auth/logout` endpoint
- [ ] Complete `GET /auth/me` endpoint
- [ ] Implement RoleGuard middleware
- [ ] Implement Product GET endpoint
- [ ] Implement Product POST endpoint

**This Week (29 Apr - 08 Mei):**
- [ ] Complete all Sprint 1 tasks
- [ ] Review [SIMPLIFIED_SPRINT_PLAN.md](../SIMPLIFIED_SPRINT_PLAN.md) daily
- [ ] Prepare for Sprint 1 Review (08 Mei)

**Next Week (11-22 Mei):**
- [ ] Start Sprint 2: Transaction Processing

---

*This is a living document. Update regularly to reflect project progress.*  
**Last Updated:** 27 April 2026
