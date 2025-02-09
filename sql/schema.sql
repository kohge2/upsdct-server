
/*MEMO アプリケーション側にデータ不整合の責任を持たせる方針にするので、制約はあまりつけない(foreign_keyなど) */
CREATE TABLE users (
    id VARCHAR(26) PRIMARY KEY,
    company_id VARCHAR(26) NOT NULL,
    email VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at DATETIME DEFAULT NULL,
    UNIQUE KEY u1 (email),
    INDEX i1 (company_id)
) DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

CREATE TABLE companies (
    id VARCHAR(26) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    representative_name VARCHAR(255) NOT NULL,
    tel VARCHAR(255) NOT NULL,
    postal_code VARCHAR(255) NOT NULL,
    address VARCHAR(255) NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at DATETIME DEFAULT NULL,
    UNIQUE KEY u1 (name)
) DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

CREATE TABLE partner_companies (
    id VARCHAR(26) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    representative_name VARCHAR(255) NOT NULL,
    tel VARCHAR(255) NOT NULL,
    postal_code VARCHAR(255) NOT NULL,
    address VARCHAR(255) NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at DATETIME DEFAULT NULL,
    UNIQUE KEY u1 (name)
) DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

CREATE TABLE partner_company_bank_account (
    id VARCHAR(26) PRIMARY KEY,
    partner_company_id VARCHAR(26) NOT NULL,
    bank_name VARCHAR(255) NOT NULL,
    branch_name VARCHAR(255) NOT NULL,
    account_type VARCHAR(255) NOT NULL,
    account_number VARCHAR(255) NOT NULL,
    account_holder_name VARCHAR(255) NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at DATETIME DEFAULT NULL,
    INDEX i1 (partner_company_id)
) DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

CREATE TABLE invoices (
    id VARCHAR(26) PRIMARY KEY,
    company_id VARCHAR(26) NOT NULL,
    partner_company_id VARCHAR(26) NOT NULL,
    published_date DATE NOT NULL,
    payment_due_date DATETIME NOT NULL,
    commission INT unsigned NOT NULL,
    commission_rate DECIMAL(6, 5) NOT NULL,
    paid_ammount INT unsigned NOT NULL,
    billed_ammount INT unsigned NOT NULL,
    invoice_status VARCHAR(255) NOT NULL,
    created_by VARCHAR(26) NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at DATETIME DEFAULT NULL,
    INDEX i1 (company_id, payment_due_date),
    INDEX i2 (partner_company_id, payment_due_date)
) DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;