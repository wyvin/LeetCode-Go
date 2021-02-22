-- 竟然考sql，出其不意。
select FirstName, LastName, City, State from Person left join Address on Person.PersonId=Address.PersonId;