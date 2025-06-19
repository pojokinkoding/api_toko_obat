-- Tabel Kategori Obat
CREATE TABLE categories (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    description TEXT,
    created_by VARCHAR(128),
    updated_by VARCHAR(128),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Tabel Produsen / Merek
CREATE TABLE manufacturers (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(150) NOT NULL,
    address TEXT,
    contact VARCHAR(100),
    created_by VARCHAR(128),
    updated_by VARCHAR(128),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Tabel Obat
CREATE TABLE medicines (
    id BIGSERIAL PRIMARY KEY,
    medicine_code VARCHAR(50) UNIQUE NOT NULL,
    name VARCHAR(255) NOT NULL,
    category_id BIGINT REFERENCES categories(id) ON DELETE SET NULL,
    manufacturer_id BIGINT REFERENCES manufacturers(id) ON DELETE SET NULL,
    type VARCHAR(50),
    description TEXT,
    price DECIMAL(10,2) NOT NULL DEFAULT 0,
    stock INTEGER NOT NULL DEFAULT 0,
    unit VARCHAR(50),
    expiry_date DATE,
    created_by VARCHAR(128),
    updated_by VARCHAR(128),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
