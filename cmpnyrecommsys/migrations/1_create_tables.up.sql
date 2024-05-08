CREATE TABLE companies (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    industry TEXT NOT NULL,
    city TEXT NOT NULL,
    createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE presets (
    id SERIAL PRIMARY KEY,
    type TEXT NOT NULL,
    companyId INTEGER NOT NULL,
    FOREIGN KEY (companyId) REFERENCES companies(id),
    revenue FLOAT NOT NULL,
    cogs FLOAT NOT NULL,
    depreciation FLOAT NOT NULL,
    longTermAssets FLOAT NOT NULL,
    shortTermAssets FLOAT NOT NULL,
    longTermLiability FLOAT NOT NULL,
    shortTermLiability FLOAT NOT NULL,
    operatingExpense FLOAT NOT NULL,
    retainedEarnings FLOAT NOT NULL,
    yearsInBusiness INTEGER NOT NULL,
    createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);