Anotações de boas praticas:

- Transactions: Chamadas de diferentes reposiitorys não devem ser injetadas, devem ser inicializadas no contexto da propria função.
- DB: Sempre passar contexto em querys no banco de dados (Uma hora vai ser uti)

QUICK COMMAND:

migrate -source file://migrations -database postgresql://postgres:123@localhost:5432/control_ssh?sslmode=disable

# **Gerenciando Migrations com `migrate`**

Boas práticas para criar, aplicar e reverter migrations sem problemas.

---

## **📋 Comandos Básicos**

### **1. Criar uma Nova Migration**

```bash
migrate create -dir migrations -ext sql NOME_DA_MIGRATION
```

- Gera arquivos `.up.sql` (aplicar) e `.down.sql` (reverter).
- Exemplo: `0001_create_users_table.up.sql`.

### **2. Aplicar Todas as Migrations Pendentes**

```bash
migrate -source file://migrations -database "postgres://user:senha@host:porta/banco?sslmode=disable" up
```

- Substitua `user`, `senha`, `host`, `porta` e `banco` pelos seus dados.

### **3. Reverter Migrations**

```bash
migrate -dir migrations -ext sql down [N]
```

- `N` = Quantidade de migrations a reverter (ex: `1` para a última).
- **Exemplo:**
  ```bash
  migrate -source file://migrations -database "postgres://..." down 1
  ```

---

## **⚠️ Boas Práticas**

### **🔹 Nunca Delete Migrations**

- Mesmo que uma migration seja revertida, **mantenha os arquivos** para evitar inconsistências.
- Se não for mais útil, renomeie para `OBSOLETE_nome_migration.sql`.

### **🔹 Sempre Teste o `.down.sql`**

- Certifique-se de que a reversão funciona antes de aplicar em produção.

### **🔹 Use Controle de Versão (Git)**

- Commite todas as migrations para evitar perdas.

### **🔹 Em Produção, Tenha Backup**

- Antes de rodar `down`, faça backup do banco:
  ```bash
  pg_dump -U user -d banco > backup.sql
  ```

---

## **🛠️ Solução de Problemas**

### **Erro: "no migration found for version X"**

- **Causa:** O banco registra uma migration que não existe mais no diretório.
- **Solução:**
  ```bash
  migrate -source file://migrations -database "postgres://..." force X
  ```
  (Substitua `X` pela versão correta.)

---

## **📌 Resumo**

| Ação                | Comando                                                                |
| ------------------- | ---------------------------------------------------------------------- |
| Criar migration     | `migrate create -dir migrations -ext sql NOME`                         |
| Aplicar (`up`)      | `migrate -source file://migrations -database "postgres://..." up`      |
| Reverter (`down N`) | `migrate -source file://migrations -database "postgres://..." down 1`  |
| Forçar versão       | `migrate -source file://migrations -database "postgres://..." force X` |

**Mantenha o histórico limpo e consistente!** 🚀

# back-ssh-control
