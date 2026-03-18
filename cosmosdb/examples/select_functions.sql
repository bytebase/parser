SELECT CONCAT(c.name, ' - ', c.country) AS label, STRINGTONUMBER(c.population) / 1000 AS populationInThousands FROM c
