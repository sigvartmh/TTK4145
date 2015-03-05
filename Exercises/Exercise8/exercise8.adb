with Ada.Text_IO, Ada.Integer_Text_IO, Ada.Numerics.Float_Random;
use  Ada.Text_IO, Ada.Integer_Text_IO, Ada.Numerics.Float_Random;

procedure exercise8 is

    Count_Failed    : exception;    -- Exception to be raised when counting fails
    Gen             : Generator;    -- Random number generator

    protected type Transaction_Manager (N : Positive) is
        entry Finished;
        function Commit return Boolean;
        procedure Signal_Abort;
    private
        Finished_Gate_Open  : Boolean := False;
        Aborted             : Boolean := False;
    
    end Transaction_Manager;
    protected body Transaction_Manager is
        entry Finished when Finished_Gate_Open or Finished'Count = N is
        begin
            ------------------------------------------
            -- PART 3: Complete the exit protocol here
            ------------------------------------------
            if Finished'Count=N-1 then
                Finished_Gate_Open:= true;
            end if;
            if Finished'Count=0 then
                Finished_Gate_Open:=false;
                Aborted:=false;
            end if;
        end Finished;

        procedure Signal_Abort is
        begin
            Aborted := True;
        end Signal_Abort;
        
        entry Wait_Until_Aborted when Aborted is
        begin
            if Wait_Until_Aborted'Count=0 then
                Aborted = false;
            end if
        end Wait_Until_Aborted;
        
    end Transaction_Manager;



    
    function Unreliable_Slow_Add (x : Integer) return Integer is
    Error_Rate : Constant := 0.15;  -- (between 0 and 1)
    Num : integer := 0;
    begin
        -------------------------------------------
        -- PART 1: Create the transaction work here
        -------------------------------------------
        case Random(Gen)>Error_Rate is
            when true => 
                delay Duration(Random(Gen)*4.4);
                Num:=x+10;
            when false =>
                delay Duration(Random(gen)*0.4);
                raise Count_Failed; --with "Unreliable_slow_Add_failed"
         end case;
         return Num;
    end Unreliable_Slow_Add;




    task type Transaction_Worker (Initial : Integer; Manager : access Transaction_Manager);
    task body Transaction_Worker is
        Num         : Integer   := Initial;
        Prev        : Integer   := Num;
        Round_Num   : Integer   := 0;
    begin
        Put_Line ("Worker" & Integer'Image(Initial) & " started");

        loop
            Put_Line ("Worker" & Integer'Image(Initial) & " started round" & Integer'Image(Round_Num));
            Round_Num := Round_Num + 1;

            ---------------------------------------
            -- PART 2: Do the transaction work here             
            ---------------------------------------
            select
                Manager.Wait_Until_Aborted;
                Num:=Prev+5;
                Put_Line("Worker " &Integer'Image(Initial) & "forward error recovery from" &IntegerImage(Prev) & "  to" & Integer'Image(Num));
            then abort
                begin
                    Num:=Unreliable_Slow_Add(Prev);
                exception
                    when Count_Failed=>
                        Manager.Signal_Abort;
                end;
            end select;
            Manager.Finished;
            Put_Line("Worker " &Integer'Image(Initial) & "comitting" & Integer'Image(Num));

            --if Manager.Commit = True then
            --    Put_Line ("  Worker" & Integer'Image(Initial) & " comitting" & Integer'Image(Num));
            --else
            --    Put_Line ("  Worker" & Integer'Image(Initial) &
            --                 " reverting from" & Integer'Image(Num) &
            --                 " to" & Integer'Image(Prev));
                -------------------------------------------
                -- PART 2: Roll back to previous value here
                -------------------------------------------
            --    Num := Prev;
            --end if;

            Prev := Num;
            delay 0.5;

        end loop;
    end Transaction_Worker;

    Manager : aliased Transaction_Manager (3);

    Worker_1 : Transaction_Worker (0, Manager'Access);
    Worker_2 : Transaction_Worker (1, Manager'Access);
    Worker_3 : Transaction_Worker (2, Manager'Access);

begin
    Reset(Gen); -- Seed the random number generator
end exercise8;



