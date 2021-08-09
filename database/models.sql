CREATE TABLE IF NOT EXISTS Medicamento (
    id_Med serial NOT NULL,
    nombre VARCHAR(150) NOT NULL,
    precio double precision NOT NULL,
    ubicacion VARCHAR(150) NOT NULL,    
    CONSTRAINT pk_IdMed PRIMARY KEY(id_Med)
);

CREATE TABLE IF NOT EXISTS Promocion (
    id_Promo integer NOT NULL,
    descripcion VARCHAR(150) NOT NULL,
    porcentaje double precision NOT NULL,
    fecha_inicio date,
    fecha_fin  date,    
    CONSTRAINT pk_IdPromo PRIMARY KEY(id_Promo)
);

CREATE TABLE IF NOT EXISTS Factura (
    id_Fac integer NOT NULL,
    med_id integer NOT NULL,
    promo_id integer NOT NULL,
    feccha_crear date,
    pago_total double precision,
    CONSTRAINT pk_IdFactura PRIMARY KEY(id_Fac),
    CONSTRAINT fk_fact_Med FOREIGN KEY(med_id) REFERENCES Medicamento(id_Med),
    CONSTRAINT fk_fact_Promo FOREIGN KEY(promo_id) REFERENCES Promocion(id_Promo)

);