CREATE TABLE armas (
    equipamento_id INTEGER PRIMARY KEY,

    tipo_arma TEXT NOT NULL
        CHECK (
            tipo_arma IN (
                'CORPO_A_CORPO',
                'DISTANCIA'
            )
        ),

    dano_base TEXT NOT NULL,

    tipo_dano TEXT NOT NULL,

    critico_margem INTEGER NOT NULL
        DEFAULT 20
        CHECK (
            critico_margem BETWEEN 1 AND 20
        ),

    critico_multiplicador INTEGER NOT NULL
        DEFAULT 2
        CHECK (
            critico_multiplicador >= 2
        ),

    alcance TEXT NOT NULL
        CHECK (
            alcance IN (
                'CORPO_A_CORPO',
                'CURTO',
                'MEDIO',
                'LONGO',
                'EXTREMO'
            )
        ),

    empunhadura TEXT NOT NULL
        CHECK (
            empunhadura IN (
                'LEVE',
                'UMA_MAO',
                'DUAS_MAOS'
            )
        ),

    recarga TEXT,

    FOREIGN KEY (equipamento_id)
        REFERENCES equipamentos(id)
        ON DELETE CASCADE
);

CREATE TABLE protecoes (
    equipamento_id INTEGER PRIMARY KEY,

    tipo_protecao TEXT NOT NULL
        CHECK (
            tipo_protecao IN (
                'LEVE',
                'PESADA',
                'ESCUDO',
                'OUTRA'
            )
        ),

    bonus_defesa INTEGER NOT NULL
        DEFAULT 0,

    penalidade_testes INTEGER NOT NULL
        DEFAULT 0
        CHECK (
            penalidade_testes <= 0
        ),

    FOREIGN KEY (equipamento_id)
        REFERENCES equipamentos(id)
        ON DELETE CASCADE
);

CREATE TABLE municoes (
    equipamento_id INTEGER PRIMARY KEY,

    duracao_quantidade INTEGER
        CHECK (
            duracao_quantidade IS NULL
            OR duracao_quantidade > 0
        ),

    duracao_unidade TEXT NOT NULL
        CHECK (
            duracao_unidade IN (
                'CENA',
                'DISPARO',
                'USO',
                'MISSAO',
                'ESPECIAL'
            )
        ),

    consumivel INTEGER NOT NULL
        DEFAULT 1
        CHECK (
            consumivel IN (0, 1)
        ),

    FOREIGN KEY (equipamento_id)
        REFERENCES equipamentos(id)
        ON DELETE CASCADE
);

CREATE TABLE arma_municoes (
    arma_equipamento_id INTEGER NOT NULL,
    municao_equipamento_id INTEGER NOT NULL,

    PRIMARY KEY (
        arma_equipamento_id,
        municao_equipamento_id
    ),

    FOREIGN KEY (arma_equipamento_id)
        REFERENCES armas(equipamento_id)
        ON DELETE CASCADE,

    FOREIGN KEY (municao_equipamento_id)
        REFERENCES municoes(equipamento_id)
        ON DELETE CASCADE,

    CHECK (
        arma_equipamento_id
        <> municao_equipamento_id
    )
);

CREATE INDEX idx_armas_tipo
    ON armas(tipo_arma);

CREATE INDEX idx_armas_alcance
    ON armas(alcance);

CREATE INDEX idx_armas_tipo_dano
    ON armas(tipo_dano);

CREATE INDEX idx_protecoes_tipo
    ON protecoes(tipo_protecao);

CREATE INDEX idx_arma_municoes_municao
    ON arma_municoes(municao_equipamento_id);


CREATE TRIGGER validar_tipo_arma_insert
BEFORE INSERT ON armas
FOR EACH ROW
WHEN (
    SELECT tipo
    FROM equipamentos
    WHERE id = NEW.equipamento_id
) <> 'ARMA'
BEGIN
    SELECT RAISE(
        ABORT,
        'o equipamento precisa ser do tipo ARMA'
    );
END;

CREATE TRIGGER validar_tipo_protecao_insert
BEFORE INSERT ON protecoes
FOR EACH ROW
WHEN (
    SELECT tipo
    FROM equipamentos
    WHERE id = NEW.equipamento_id
) <> 'PROTECAO'
BEGIN
    SELECT RAISE(
        ABORT,
        'o equipamento precisa ser do tipo PROTECAO'
    );
END;

CREATE TRIGGER validar_tipo_municao_insert
BEFORE INSERT ON municoes
FOR EACH ROW
WHEN (
    SELECT tipo
    FROM equipamentos
    WHERE id = NEW.equipamento_id
) <> 'MUNICAO'
BEGIN
    SELECT RAISE(
        ABORT,
        'o equipamento precisa ser do tipo MUNICAO'
    );
END;