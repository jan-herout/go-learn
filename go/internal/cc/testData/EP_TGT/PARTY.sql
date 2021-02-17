
/*==============================================================*/
/* Table: PARTY                                                 */
/*==============================================================*/
create multiset table PARTY (
    party_id                        INTEGER                        not null,
    party_cd                        VARCHAR(100)                   not null   character set unicode,
    party_type_id                   SMALLINT                       not null,
    src_cls_id                      SMALLINT                       not null,
    tt_start_dttm                   TIMESTAMP(6) WITH TIME ZONE    not null,
    tt_end_dttm                     TIMESTAMP(6) WITH TIME ZONE    not null,
    PERIOD FOR tt_per                          (tt_start_dttm, tt_end_dttm) AS TRANSACTIONTIME,
    load_id                         INTEGER                        not null,
    load_dttm                       TIMESTAMP(0)                   not null,
    map_id                          INTEGER                        not null,
    hist_type                       INTEGER                        not null
)
primary index PARTY_NUPI (
    party_id
 );

comment on table PARTY is 'A party is any individual, organization or household that is of interest to the enterprise';
comment on column PARTY.party_id is  'An unique identifier of Party';
comment on column PARTY.party_cd is  'A code of Party';
comment on column PARTY.party_type_id is  'An unique identifier of Party Type';
comment on column PARTY.src_cls_id is  'An unique identifier of Source Class';
comment on column PARTY.tt_start_dttm is  'Start of transactiontime validity (temporal)';
comment on column PARTY.tt_end_dttm is  'End of transactiontime validity (temporal)';
comment on column PARTY.tt_per is  'Transaction validity - Teradata TDF';
comment on column PARTY.load_id is  'Audit column - load instance ID';
comment on column PARTY.load_dttm is  'Audit column - load date when the record was loaded';
comment on column PARTY.map_id is  'Audit column - mapping which loaded the record';
comment on column PARTY.hist_type is  'Audit column - hist type';
	