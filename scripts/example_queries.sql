-- Получить все курсы валют
SELECT * FROM exchange_rates;

-- Получить курсы за конкретную дату
SELECT * FROM exchange_rates 
WHERE Date = '2024-03-14';

-- Получить курс конкретной валюты
SELECT * FROM exchange_rates 
WHERE Cur_Abbreviation = 'USD' 
ORDER BY Date DESC;

-- Получить последние курсы всех валют
SELECT r.* 
FROM exchange_rates r
INNER JOIN (
    SELECT Cur_Abbreviation, MAX(Date) as MaxDate 
    FROM exchange_rates 
    GROUP BY Cur_Abbreviation
) latest 
ON r.Cur_Abbreviation = latest.Cur_Abbreviation 
AND r.Date = latest.MaxDate;


SELECT 
    r1.Date,
    r1.Cur_Abbreviation,
    r1.Cur_OfficialRate as CurrentRate,
    r2.Cur_OfficialRate as PreviousRate,
    (r1.Cur_OfficialRate - r2.Cur_OfficialRate) as Difference
FROM exchange_rates r1
LEFT JOIN exchange_rates r2 
    ON r1.Cur_Abbreviation = r2.Cur_Abbreviation
    AND r1.Date = DATE_ADD(r2.Date, INTERVAL 1 DAY)
ORDER BY r1.Date DESC, r1.Cur_Abbreviation;
