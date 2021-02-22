-- distinct 用于去重
-- 没想到直接select as就可以解决返回null的问题。
SELECT
    (SELECT DISTINCT Salary
    FROM Employee
    ORDER BY  Salary DESC limit 1 offset 1) AS SecondHighestSalary;