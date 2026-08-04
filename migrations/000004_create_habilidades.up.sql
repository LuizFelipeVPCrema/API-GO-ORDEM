CREATE TABLE habilidades (
    id INTEGER PRIMARY KEY AUTOINCREMENT,

    codigo TEXT NOT NULL UNIQUE,
    nome TEXT NOT NULL,

    categoria TEXT NOT NULL
        CHECK (
            categoria IN (
                'HABILIDADE_CLASSE',
                'PODER_CLASSE',
                'HABILIDADE_ORIGEM',
                'HABILIDADE_TRILHA',
                'PODER_PARANORMAL',
                'PODER_GERAL',
                'OUTRA'
            )
        ),

    tipo_ativacao TEXT NOT NULL
        CHECK (
            tipo_ativacao IN (
                'PASSIVA',
                'LIVRE',
                'MOVIMENTO',
                'PADRAO',
                'COMPLETA',
                'REACAO',
                'ESPECIAL'
            )
        ),

    custo_pe_base INTEGER
        CHECK (
            custo_pe_base IS NULL
            OR custo_pe_base >= 0
        ),

    custo_pe_variavel INTEGER NOT NULL DEFAULT 0
        CHECK (custo_pe_variavel IN (0, 1)),

    descricao_resumida TEXT,

    fonte_regra TEXT NOT NULL,
    versao_regra TEXT,

    ativa INTEGER NOT NULL DEFAULT 1
        CHECK (ativa IN (0, 1)),

    criado_em DATETIME NOT NULL
        DEFAULT CURRENT_TIMESTAMP,

    atualizado_em DATETIME NOT NULL
        DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE classe_habilidades (
    id INTEGER PRIMARY KEY AUTOINCREMENT,

    classe_id INTEGER NOT NULL,
    habilidade_id INTEGER NOT NULL,

    nex_minimo INTEGER NOT NULL DEFAULT 0
        CHECK (
            nex_minimo >= 0
            AND nex_minimo <= 100
        ),

    forma_aquisicao TEXT NOT NULL
        CHECK (
            forma_aquisicao IN (
                'AUTOMATICA',
                'ESCOLHA'
            )
        ),

    ordem_exibicao INTEGER NOT NULL DEFAULT 0,

    UNIQUE (
        classe_id,
        habilidade_id,
        nex_minimo
    ),

    FOREIGN KEY (classe_id)
        REFERENCES classes(id)
        ON DELETE CASCADE,

    FOREIGN KEY (habilidade_id)
        REFERENCES habilidades(id)
        ON DELETE CASCADE
);

CREATE INDEX idx_classe_habilidades_classe
    ON classe_habilidades(classe_id);

CREATE INDEX idx_classe_habilidades_habilidade
    ON classe_habilidades(habilidade_id);

CREATE INDEX idx_classe_habilidades_nex
    ON classe_habilidades(classe_id, nex_minimo);