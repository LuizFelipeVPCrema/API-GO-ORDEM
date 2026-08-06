CREATE TABLE personagens (
    id INTEGER PRIMARY KEY AUTOINCREMENT,

    nome TEXT NOT NULL
        CHECK (LENGTH(TRIM(nome)) > 0),

    jogador TEXT,

    classe_id INTEGER NOT NULL,

    nex INTEGER NOT NULL DEFAULT 0
        CHECK (
            nex >= 0
            AND nex <= 100
        ),

    pontos_prestigio INTEGER NOT NULL DEFAULT 0
        CHECK (pontos_prestigio >= 0),

    idade INTEGER
        CHECK (
            idade IS NULL
            OR idade > 0
        ),

    aparencia TEXT,
    personalidade TEXT,
    historia TEXT,
    objetivo TEXT,

    ativa INTEGER NOT NULL DEFAULT 1
        CHECK (ativa IN (0, 1)),

    criado_em DATETIME NOT NULL
        DEFAULT CURRENT_TIMESTAMP,

    atualizado_em DATETIME NOT NULL
        DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (classe_id)
        REFERENCES classes(id)
        ON DELETE RESTRICT
);

CREATE TABLE personagem_atributos (
    personagem_id INTEGER PRIMARY KEY,

    agilidade INTEGER NOT NULL DEFAULT 0
        CHECK (agilidade >= 0),

    forca INTEGER NOT NULL DEFAULT 0
        CHECK (forca >= 0),

    intelecto INTEGER NOT NULL DEFAULT 0
        CHECK (intelecto >= 0),

    presenca INTEGER NOT NULL DEFAULT 0
        CHECK (presenca >= 0),

    vigor INTEGER NOT NULL DEFAULT 0
        CHECK (vigor >= 0),

    atualizado_em DATETIME NOT NULL
        DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (personagem_id)
        REFERENCES personagens(id)
        ON DELETE CASCADE
);

CREATE TABLE personagem_recursos (
    personagem_id INTEGER PRIMARY KEY,

    pv_atual INTEGER NOT NULL DEFAULT 0
        CHECK (pv_atual >= 0),

    pv_maximo INTEGER NOT NULL DEFAULT 0
        CHECK (pv_maximo >= 0),

    pe_atual INTEGER NOT NULL DEFAULT 0
        CHECK (pe_atual >= 0),

    pe_maximo INTEGER NOT NULL DEFAULT 0
        CHECK (pe_maximo >= 0),

    sanidade_atual INTEGER NOT NULL DEFAULT 0
        CHECK (sanidade_atual >= 0),

    sanidade_maxima INTEGER NOT NULL DEFAULT 0
        CHECK (sanidade_maxima >= 0),

    atualizado_em DATETIME NOT NULL
        DEFAULT CURRENT_TIMESTAMP,

    CHECK (pv_atual <= pv_maximo),
    CHECK (pe_atual <= pe_maximo),
    CHECK (sanidade_atual <= sanidade_maxima),

    FOREIGN KEY (personagem_id)
        REFERENCES personagens(id)
        ON DELETE CASCADE
);

CREATE TABLE personagem_pericias (
    id INTEGER PRIMARY KEY AUTOINCREMENT,

    personagem_id INTEGER NOT NULL,
    pericia_id INTEGER NOT NULL,

    grau_treinamento TEXT NOT NULL DEFAULT 'DESTREINADO'
        CHECK (
            grau_treinamento IN (
                'DESTREINADO',
                'TREINADO',
                'VETERANO',
                'EXPERT'
            )
        ),

    bonus_outros INTEGER NOT NULL DEFAULT 0,

    especialidade TEXT NOT NULL DEFAULT '',

    criado_em DATETIME NOT NULL
        DEFAULT CURRENT_TIMESTAMP,

    atualizado_em DATETIME NOT NULL
        DEFAULT CURRENT_TIMESTAMP,

    UNIQUE (
        personagem_id,
        pericia_id,
        especialidade
    ),

    FOREIGN KEY (personagem_id)
        REFERENCES personagens(id)
        ON DELETE CASCADE,

    FOREIGN KEY (pericia_id)
        REFERENCES pericias(id)
        ON DELETE RESTRICT
);

CREATE TABLE personagem_habilidades (
    id INTEGER PRIMARY KEY AUTOINCREMENT,

    personagem_id INTEGER NOT NULL,
    habilidade_id INTEGER NOT NULL,

    fonte_aquisicao TEXT NOT NULL
        CHECK (
            fonte_aquisicao IN (
                'CLASSE',
                'ORIGEM',
                'TRILHA',
                'PODER_PARANORMAL',
                'OUTRA'
            )
        ),

    adquirida_em_nex INTEGER
        CHECK (
            adquirida_em_nex IS NULL
            OR (
                adquirida_em_nex >= 0
                AND adquirida_em_nex <= 100
            )
        ),

    observacao TEXT,

    ativa INTEGER NOT NULL DEFAULT 1
        CHECK (ativa IN (0, 1)),

    criado_em DATETIME NOT NULL
        DEFAULT CURRENT_TIMESTAMP,

    UNIQUE (
        personagem_id,
        habilidade_id
    ),

    FOREIGN KEY (personagem_id)
        REFERENCES personagens(id)
        ON DELETE CASCADE,

    FOREIGN KEY (habilidade_id)
        REFERENCES habilidades(id)
        ON DELETE RESTRICT
);

CREATE TABLE personagem_rituais (
    id INTEGER PRIMARY KEY AUTOINCREMENT,

    personagem_id INTEGER NOT NULL,
    ritual_id INTEGER NOT NULL,

    fonte_aquisicao TEXT NOT NULL
        CHECK (
            fonte_aquisicao IN (
                'CLASSE',
                'HABILIDADE',
                'PODER_PARANORMAL',
                'OUTRA'
            )
        ),

    adquirido_em_nex INTEGER
        CHECK (
            adquirido_em_nex IS NULL
            OR (
                adquirido_em_nex >= 0
                AND adquirido_em_nex <= 100
            )
        ),

    habilitado INTEGER NOT NULL DEFAULT 1
        CHECK (habilitado IN (0, 1)),

    observacao TEXT,

    criado_em DATETIME NOT NULL
        DEFAULT CURRENT_TIMESTAMP,

    UNIQUE (
        personagem_id,
        ritual_id
    ),

    FOREIGN KEY (personagem_id)
        REFERENCES personagens(id)
        ON DELETE CASCADE,

    FOREIGN KEY (ritual_id)
        REFERENCES rituais(id)
        ON DELETE RESTRICT
);

CREATE INDEX idx_personagens_nome
    ON personagens(nome);

CREATE INDEX idx_personagens_classe
    ON personagens(classe_id);

CREATE INDEX idx_personagens_ativa
    ON personagens(ativa);

CREATE INDEX idx_personagem_pericias_personagem
    ON personagem_pericias(personagem_id);

CREATE INDEX idx_personagem_habilidades_personagem
    ON personagem_habilidades(personagem_id);

CREATE INDEX idx_personagem_rituais_personagem
    ON personagem_rituais(personagem_id);