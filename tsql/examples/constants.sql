create procedure sp (@val numeric) as begin print @val
end
go
    exec sp $ 5
go
    exec sp $ -5
go
    exec sp $ + 5
go
    exec sp - $ 5
go
    exec sp - $ -5
go
    exec sp -1
go
select
    $ 5
go
select
    $ -5
go
select
    $ + 5
go
select
    - $ 5
go
select
    - $ -5
go
select
    + $ -5
go
select
    -1
go
select
    - -1
go
select
    + 1
go