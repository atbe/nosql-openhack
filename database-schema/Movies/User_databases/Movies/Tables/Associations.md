#### 

[Movies](../index.md) > [Tables](Tables.md) > dbo.Associations

# ![Tables](../../../Images/Table32.png) [dbo].[Associations]

---

## <a name="#description"></a>MS_Description

Experimental: Used for associating movies for recommendations. Generated by custom ML model.

## <a name="#columns"></a>Columns

| Key | Name | Data Type | Max Length (Bytes) | Nullability |
|---|---|---|---|---|
| [![Cluster Primary Key PK_Associations: id](../../../Images/pkcluster.png)](#indexes) | id | varchar(100) | 100 | NOT NULL |
|  | created | datetime | 8 | NULL allowed |
|  | confidence | float | 8 | NULL allowed |
|  | support | float | 8 | NULL allowed |
|  | source | int | 4 | NOT NULL |
|  | target | int | 4 | NOT NULL |


---

## <a name="#sqlscript"></a>SQL Script

```sql
CREATE TABLE [dbo].[Associations]
(
[id] [varchar] (100) COLLATE SQL_Latin1_General_CP1_CI_AS NOT NULL,
[created] [datetime] NULL,
[confidence] [float] NULL,
[support] [float] NULL,
[source] [int] NOT NULL,
[target] [int] NOT NULL
)
GO
ALTER TABLE [dbo].[Associations] ADD CONSTRAINT [PK_Associations] PRIMARY KEY CLUSTERED  ([id])
GO
EXEC sp_addextendedproperty N'MS_Description', N'Experimental: Used for associating movies for recommendations. Generated by custom ML model.', 'SCHEMA', N'dbo', 'TABLE', N'Associations', NULL, NULL
GO

```


---

###### Author:  Contoso Movies, Ltd.

###### Copyright 2019 - All Rights Reserved

###### Created: 2019/11/27
