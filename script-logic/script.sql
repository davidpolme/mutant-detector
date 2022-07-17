declare

v1 varchar2(1);
v2 varchar2(1);
v3 varchar2(1);

begin
    v1 := 'a';
    v2 := 'b';
    v3 := v1 == v2;
    dbms_output.put_line(v1 || v2 || v3);
    end;
end;