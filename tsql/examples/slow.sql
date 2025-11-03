

if exists (select * from sysobjects where name = 'up_com_test')
  drop proc up_com_test
go

create proc up_com_test
  @p_vipa1111111111111  dtint,
  @p_tblname      dtchar32

 -- 
 as
 set nocount on

 declare @a1111111111111   dtint,
 		 @sysdate    dtdate,
 		 @a2222222222  dtdate,
 		 @a333333333333     dtkind,
 		 @a444444444444444   dtint,
 		 @vipid      dtchar32,
 		 @vipa333333333333  dtkind
  
         
 exec nb_com_getsysconfig @a1111111111111 output, @sysdate output , @a2222222222 output, @a333333333333 output
 select @a444444444444444 = paravalue from dbo.mixedconfig where a1111111111111 = @a1111111111111 and paraid = 'vipsubsys_sn'
 select @vipid = vipid,@vipa333333333333 = a333333333333 from dbo.vipconfig where a1111111111111 = @a1111111111111 and vip_a1111111111111 = @p_vipa1111111111111
 if @a333333333333 !='0'
 begin
  select  errorcode = -1, errormsg = '执行!'
  return  -1
 end
 
 if @vipa333333333333 !='0'
 begin
  select  errorcode = -1, errormsg = '准备!'
  return  -1
 end


 declare @gmhardware dtkind
 set @gmhardware = '0'
 select @gmhardware = paravalue from dbo.mixedconfig where a1111111111111 = @a1111111111111 and paraid = 'gmhardware'


 create table #stktype
 (a1111111111111  int default 0,
  market     char(1) default ' ',
  stktype    char(1) default ' ',
  feestktype char(1)  default ' ',
  moneytype  char(1)  default ' '
  )

  create table #diffstktype
  (market    char(1)   default ' ',
   stktype    char(1)   default ' ')
 
 
   create table #bizConvert
   (fstkbiz          int  default  0,
    tstkbiz          int   default 0)
  
 if @p_tblname = 'ITF_USERS'
 begin
    delete from dbo.ITF_USERS where SYNC_TYPE = '1' and vipa1111111111111 = @p_vipa1111111111111  
    insert into dbo.ITF_USERS
       (USER_CODE, USER_ROLES, USER_NAME, USER_TYPE, INT_ORG, 
        ID_TYPE, ID_CODE, OPEN_DATE, vipa1111111111111, SYNC_TYPE       )
    select a.custid                      as USER_CODE,   
           '1'                           as USER_ROLES,   
       ISNULL(rtrim(a.custname),' ')     as USER_NAME,     
       b.singleflag                      as USER_TYPE,    
       a.orgid                           as INT_ORG,       
       dbo.fn_ksts_idtype2ID_TYPE(b.idtype)  as ID_TYPE,     

       case when @gmhardware = '2' then b.sm_data1 else b.idno end as ID_CODE,
       b.opendate                        as OPEN_DATE,    
       @p_vipa1111111111111                    as vipa1111111111111, 
       '1'                               as SYNC_TYPE     
    from dbo.customer a 
    inner join dbo.c666666666666666 b with (nolock) on a.custid = b.custid and a.a1111111111111 = b.a1111111111111
    inner join dbo.vip_customers d with (nolock) on a.a1111111111111 = d.a1111111111111 and a.custid = d.custid and a.orgid = d.orgid
    inner join dbo.fundinfo e with (nolock) on d.a1111111111111 = e.a1111111111111 and d.orgid = e.orgid and d.fundid = e.fundid and a.custid = e.custid
    where a.custprop = '0' and a.a333333333333 != '*' and e.fundtype != 'T' and e.a333333333333 != '*'
      and d.vipid = @vipid 
      and d.a333333333333 = '3'
    if @@error != 0
    begin
       select errorcode = -1 , errormsg = '失败'
    end

    
    select errorcode = 0 , errormsg = 'ITF_USERS数据处理成功'
    update dbo.itf_sync_tables set prepcnt = (select count(*) from dbo.ITF_USERS  where SYNC_TYPE = '1' and vipa1111111111111 = @p_vipa1111111111111),a333333333333 = '1'
                where vipid = @vipid and tablename = @p_tblname and direct = 'e'
    select tablecnt = count(*) from dbo.ITF_USERS  where SYNC_TYPE = '1' and vipa1111111111111 = @p_vipa1111111111111
 end
  
 else if @p_tblname = 'ITF_CUSTOMER'
 begin
    delete from dbo.ITF_CUSTOMER  where SYNC_TYPE = '1' and vipa1111111111111 = @p_vipa1111111111111 
    insert into dbo.ITF_CUSTOMER
    (CUST_CODE, CUST_TYPE, CUST_CLS, CUST_a333333333333, CARD_ID,CRITERION, RISK_FACTOR,INT_ORG,OPEN_DATE, 
     CHANNELS,REMARK,CREDIT_LVL,REMOTE_PROTOCOL,CUST_SOURCE,SERVICE_LVL,CUST_EXT_ATTR,CLOSE_DATE,
     vipa1111111111111, SYNC_TYPE)
    select 
	   a.custid                                      as CUST_CODE,   
	   a.custprop                                    as USER_TYPE,      
	   a.custkind                                    as CUST_CLS,      
	   a.a333333333333                                      as CUST_a333333333333,   
	   case when len(ltrim(rtrim(a.custcard))) < 1 then ' ' 
	          else ltrim(rtrim(a.custcard)) 
	   end                                           as CARD_ID,     
	   case when len(b.criterion) < 1 then ' ' 
	        else b.criterion 
	   end                                           as CRITERION,     
	   case when len(b.riskfactor) < 1 then ' ' 
	        else b.riskfactor 
	   end                                           as RISK_FACTOR,    
	   a.orgid                       								 as INT_ORG,        
     b.opendate                    								 as OPEN_DATE,      
     left(ltrim(rtrim(ISNULL(d.operway,''))),64)   as CHANNELS, 			 
     b.remark                      								 as REMARK,         
     ''                            								 as CREDIT_LVL,      
     ''                            								 as REMOTE_PROTOCOL, 
     ''                            								 as CUST_SOURCE,      
     ''                            								 as SERVICE_LVL,     
     ''                            								 as CUST_EXT_ATTR,   
     b.closedate                   								 as CLOSE_DATE,     
     @p_vipa1111111111111                                as vipa1111111111111,     
     '1'                                           as SYNC_TYPE        
  	FROM dbo.customer a
	   inner join dbo.c666666666666666 b with (nolock) on a.custid = b.custid and a.a1111111111111 = b.a1111111111111
	   inner join dbo.vip_customers c with (nolock) on a.a1111111111111 = c.a1111111111111 and a.custid = c.custid and a.orgid = c.orgid
	   inner join dbo.fundinfo d with (nolock) on c.a1111111111111 = d.a1111111111111 and c.orgid = d.orgid and c.fundid = d.fundid and a.custid = d.custid
	WHERE a.a333333333333 != '*' and d.fundtype != 'T' and d.a333333333333 != '*' and a.custprop = '0'
	  and c.vipid = @vipid
    and c.a333333333333 = '3'
    if @@error != 0
    begin
       select errorcode = -1, errormsg = '失败'
       return -1
    end
    
    select errorcode = 0 , errormsg = '成功'
    update dbo.itf_sync_tables set prepcnt = (select count(*) from dbo.ITF_CUSTOMER where  SYNC_TYPE = '1' and vipa1111111111111 = @p_vipa1111111111111),a333333333333 = '1'
                where vipid = @vipid and tablename = @p_tblname and direct = 'e'
    select tablecnt = count(*) from dbo.ITF_CUSTOMER where  SYNC_TYPE = '1' and vipa1111111111111 = @p_vipa1111111111111
 end
  
 else if @p_tblname = 'ITF_CUST_AGREEMENT'
 begin
    delete from dbo.ITF_CUST_AGREEMENT  where SYNC_TYPE = '1' and vipa1111111111111 = @p_vipa1111111111111 
    insert into dbo.ITF_CUST_AGREEMENT
		(CUST_CODE, CUST_AGMT_TYPE, EFT_DATE, EXP_DATE, UPD_DATE, 
		 vipa1111111111111, SYNC_TYPE , CUACCT_CODE,TRDACCT,STKBD)
    select distinct
      b.custid       					as CUST_CODE,        
      dbo.fn_ksts_khqx2AGREEMENT(c.subitem)                     as CUST_AGMT_TYPE,  
      isnull(d.effectdate, 19990101)   					as EFT_DATE,                       
      case when c.recordflag = '1' then isnull(d.expiredate,19990101) else  isnull(d.expiredate, 20990101)  end  					as EXP_DATE,             
      isnull(d.procdate,   19990101)     				as UPD_DATE,       
      @p_vipa1111111111111          as vipa1111111111111,     
      '1'                     as SYNC_TYPE,      
      b.fundid                          as CUACCT_CODE,
      b.secuid                          as TRDACCT,
       dbo.fn_ksts_market2STKBD(b.market, '0') as STKBD
	from dbo.vip_customers a 
	   inner join (select distinct x.a1111111111111 a1111111111111, x.orgid orgid, x.custid custid, x.fundid, y.secuid,y.market,x.fundright+case when x.fundright!='' then ',' else '' end+y.securight fundright from dbo.fundinfo x, dbo.secuid y where x.custid = y.custid and x.orgid = y.orgid and y.a333333333333<>'*' and x.a333333333333<>'*') b
	       on a.custid = b.custid and a.orgid = b.orgid and a.a1111111111111 = b.a1111111111111  and a.fundid = b.fundid
	   inner join (select distinct rtrim(subitem) subitem , recordflag from dbo.dictvalue a, dbo.rightpropery b 
	    where dictitem = 'Bkhqx' and subitem in ('03','04','0A','0G','0H','0J','0K','0h','0V','0X','2k','0b','0e','0p','0q','0m','0v','0k','0^','0s','0N','0n','0$','0w','1e','1g','2C','2h','2i','2j','2l','1n','1q','2q','2H')--1n 基础设施基金交易权限
	      and a.subitem = b.rightid and a.a1111111111111 = b.a1111111111111
	    ) c on 1=1
	   left join (select  a1111111111111, orgid, custid, market,secuid, fundright,min(effectdate) effectdate, max(expiredate) expiredate, min(procdate) procdate from dbo.fundrightvalid where a333333333333 <> '0'
	              group by a1111111111111,orgid,custid,market,secuid,fundright ) d 
	            on a.custid = d.custid and a.orgid = d.orgid and a.a1111111111111 = d.a1111111111111 AND c.subitem = d.fundright 
	            and b.market = d.market and b.secuid = d.secuid
	where a.vipid = @vipid and a.a333333333333= '3'
	and charindex(c.subitem, b.fundright) > 0
	

  
	if @@error != 0
    begin
       select errorcode = -1, errormsg = '失败'
       return -1
    end
    

    select errorcode = 0 , errormsg = '成功'
    update dbo.itf_sync_tables set prepcnt = (select count(*) from dbo.ITF_CUST_AGREEMENT where SYNC_TYPE = '1' and vipa1111111111111 = @p_vipa1111111111111),a333333333333 = '1'
                where vipid = @vipid and tablename = @p_tblname and direct = 'e'
    select tablecnt = count(*) from dbo.ITF_CUST_AGREEMENT where SYNC_TYPE = '1' and vipa1111111111111 = @p_vipa1111111111111
 end 

  else if @p_tblname = 'ITF_CUST_STK_TRD_CTRL'
 begin
    delete from dbo.ITF_CUST_STK_TRD_CTRL where vipa1111111111111 = @p_vipa1111111111111 and SYNC_TYPE = '1' 
    insert into dbo.ITF_CUST_STK_TRD_CTRL
       (  SYNC_TYPE , INT_ORG,CUST_CODE ,CUACCT_CODE ,STKBD ,TRDACCT ,STK_BIZES ,STK_CLSES ,STK_CODE ,GRANT_TYPE ,STK_BIZ_ACTIONS,vipa1111111111111)
  select distinct 
       '1'                                       as SYNC_TYPE,
       a.orgid                                   as INT_ORG,  
       a.custid                                  as CUST_CODE,    
       a.fundid                                  as CUACCT_CODE,  
      dbo.fn_ksts_market2STKBD(a.market, '0')    as STKBD,    
       '@'                                       as TRDACCT,    
       dbo.fn_ksts_bsflag2STK_BIZ_EX(a.bsflag,a.stktype)   as STK_BIZES, 
       dbo.fn_ksts_stktype2STK_CLS(a.stktype,'0',0,a.market, ' ', ' ') as STK_CLSES,
       a.stkcode_s                                 as STK_CODE,   
       case when a.stk_direct = '0' then '1' else '0' end    as GRANT_TYPE, 
        '@'                                                  as STK_BIZ_ACTIONS,
       @p_vipa1111111111111                            as vipa1111111111111
    from dbo.fundtrdctrl a inner join dbo.vip_customers b with (nolock) 
       on a.custid = b.custid and a.fundid = b.fundid and a.a1111111111111 = b.a1111111111111
    where b.vipid = @vipid and b.a333333333333 = '3'
    and dbo.fn_ksts_bsflag2STK_BIZ_EX(a.bsflag,a.stktype) <> 0
    if @@error != 0
    begin
       select errorcode = -1 , errormsg = '失败'
    end

   select errorcode = 0 , errormsg = 'STK_HS_DETAIL数据处理成功'    
	update dbo.itf_sync_tables set prepcnt = (select count(*) from dbo.ITF_CUST_STK_TRD_CTRL where SYNC_TYPE = '1' and vipa1111111111111 = @p_vipa1111111111111),a333333333333 = '1'
                where vipid = @vipid and tablename = @p_tblname and direct = 'e'
    
    select tablecnt = count(*) from dbo.ITF_CUST_STK_TRD_CTRL where SYNC_TYPE = '1' and vipa1111111111111 = @p_vipa1111111111111
 end
 else if @p_tblname = 'ITF_CUACCT'
 begin
    delete from dbo.ITF_CUACCT  where SYNC_TYPE = '1' and vipa1111111111111 = @p_vipa1111111111111 
    insert into dbo.ITF_CUACCT
       (USER_CODE, CUACCT_CODE, CUACCT_ATTR, CUACCT_CLS, CUACCT_LVL, 
        CUACCT_GRP, CUACCT_a333333333333, MAIN_FLAG, INT_ORG, OPEN_DATE, 
        PRINT_DATE, vipa1111111111111, SYNC_TYPE, CLOSE_DATE, FUND_LIMITE)
    select 
       a.custid       													as USER_CODE,
       a.fundid       													as CUACCT_CODE,
       '0'            													as CUACCT_ATTR,
       case when len(a.fundkind) < 1 then ' ' 
            else a.fundkind
       end   																		as CUACCT_CLS,
       case when len(a.fundlevel) < 1 then ' ' 
       			else a.fundlevel 
       end 																			as CUACCT_LVL,
       case when len(a.fundgroup) < 1 then ' ' 
       	else a.fundgroup 
       end 																			as CUACCT_GRP,
       a.a333333333333               									as CUACCT_a333333333333,
       case when a.fundtype = 'T' then '0' 
       			else '1' 
       end             													as MAIN_FLAG,
       a.orgid        													as INT_ORG,
       opendate       													as OPEN_DATE,
       printdate      													as PRINT_DATE,
       @p_vipa1111111111111 													as vipa1111111111111,
       '1'           														as SYNC_TYPE,
       a.closedate                                          as CLOSE_DATE,
       a.fundlimit                                          as FUND_LIMITE
    from dbo.fundinfo a
    inner join dbo.vip_customers b with (nolock) on a.a1111111111111 = b.a1111111111111 and a.orgid = b.orgid and a.custid = b.custid and a.fundid = b.fundid
    where 1 = 1 and a.a333333333333 != '*'
      and b.vipid = @vipid
      and b.a333333333333 = '3'
    if @@error != 0
    begin
       select errorcode = -1, errormsg = '失败'
       return -1
    end
    
    select errorcode = 0 , errormsg = '成功'
    update dbo.itf_sync_tables set prepcnt = (select count(*) from dbo.ITF_CUACCT where SYNC_TYPE = '1' and vipa1111111111111 = @p_vipa1111111111111),a333333333333 = '1'
                where vipid = @vipid and tablename = @p_tblname and direct = 'e'
    select tablecnt = count(*) from dbo.ITF_CUACCT where SYNC_TYPE = '1' and vipa1111111111111 = @p_vipa1111111111111
 end 
 
 else if @p_tblname = 'ITF_STK_TRDACCT'
 begin
    declare @secusno dtsno
    set @secusno = 0
    select @secusno = secusno from dbo.seculist where a1111111111111 = @a1111111111111 and flag = '1'
   
    delete from dbo.ITF_STK_TRDACCT  where SYNC_TYPE = '1' and vipa1111111111111 = @p_vipa1111111111111 
    insert into dbo.ITF_STK_TRDACCT
       (CUST_CODE, CUACCT_CODE, INT_ORG, STKEX, STKBD, 
        TRDACCT, TRDACCT_SN, TRDACCT_EXID, TRDACCT_TYPE, TRDACCT_EXCLS, 
        TRDACCT_NAME, TRDACCT_a333333333333, TREG_a333333333333, BREG_a333333333333, STKPBU, 
        FIRMID, ID_TYPE, ID_CODE, ID_ISS_AGCY, ID_EXP_DATE, 
        OPEN_DATE, vipa1111111111111, SYNC_TYPE,TRDACCT_MAIN_FLAG
       )
    select
       a.custid               as CUST_CODE,
       b.fundid               as CUACCT_CODE,
       a.orgid                as INT_ORG,
       dbo.fn_ksts_market2STKEX(a.market) as STKEX,
       dbo.fn_ksts_market2STKBD(a.market, '0') as STKBD,
       a.secuid               as TRDACCT,
       '0'                    as TRDACCT_SN,
       case when len(a.rptsecuid) < 1 then ' ' else a.rptsecuid end     as TRDACCT_EXID,
       case when len(a.foreignflag) < 1 then ' ' else a.foreignflag end as TRDACCT_TYPE,
       case when @secusno = 24 then c.singleflag else dbo.fn_ksts_secucls2TRDACCT_EXCLS(a.secucls, c.singleflag) end as TRDACCT_EXCLS,
       case when len(ltrim(rtrim(a.name))) < 1 then ' ' else ISNULL(ltrim(rtrim(a.name)),' ') end as TRDACCT_NAME,
       a.a333333333333               as TRDACCT_a333333333333,
       dbo.fn_ksts_regflag2TREG_a333333333333(a.regflag) as TREG_a333333333333,
       dbo.fn_ksts_regflag2TREG_a333333333333(a.bondreg) as BREG_a333333333333,
       a.seat               	as STKPBU,          
       case when a.market = '3' then isnull(j.firmid, ' ') else ' ' end as FIRMID,
       dbo.fn_ksts_idtype2ID_TYPE(a.idtype) as ID_TYPE,
       case when @gmhardware = '2' then c.sm_data1 else c.idno end as ID_CODE,
       case when len(c.policeorg) < 1 then ' ' else c.policeorg end  as ID_ISS_AGCY,
       c.idenddate            as ID_EXP_DATE,
       a.opendate             as OPEN_DATE,
       @p_vipa1111111111111         as vipa1111111111111,
       '1'                    as SYNC_TYPE,
       case when a.secuseq = 0 then '1' else '0' end    as TRDACCT_MAIN_FLAG  
    from dbo.secuid a 
    inner join dbo.vip_customers b with (nolock) on a.custid = b.custid and a.a1111111111111 = b.a1111111111111 and a.orgid = b.orgid 
    inner join dbo.c666666666666666 c with (nolock) on b.custid = c.custid and b.a1111111111111 = c.a1111111111111
    inner join dbo.fundinfo x with (nolock) on b.custid = x.custid and b.fundid = x.fundid and b.orgid = x.orgid and b.a1111111111111 = x.a1111111111111 
    left join dbo.orgseat i on a.orgid = i.orgid and a.market = i.market and (i.pt_bsflag = '@' or charindex('0B', i.pt_bsflag) > 0 )
    left join dbo.seat j on a.a1111111111111 = j.a1111111111111 and a.market = j.market and i.seat = j.seat
    where a.a333333333333 != '*' and x.a333333333333 != '*' and x.fundtype != 'T'
      and b.vipid = @vipid
      and b.a333333333333 = '3' 
    if @@error != 0
    begin
       select errorcode = -1, errormsg = '失败'
       return -1
    end
    
    update dbo.ITF_STK_TRDACCT  set 
      TREG_a333333333333 = dbo.fn_ksts_regflag2TREG_a333333333333(b.regflag)
      from dbo.secuid b where TRDACCT = b.secuid and CUST_CODE = b.custid and b.market='1' and STKBD = '13'
    

    
    select errorcode = 0 , errormsg = '成功'
    update dbo.itf_sync_tables set prepcnt = (select count(*) from dbo.ITF_STK_TRDACCT where SYNC_TYPE = '1' and vipa1111111111111 = @p_vipa1111111111111),a333333333333 = '1'
                where vipid = @vipid and tablename = @p_tblname and direct = 'e'
    select tablecnt = count(*) from dbo.ITF_STK_TRDACCT where SYNC_TYPE = '1' and vipa1111111111111 = @p_vipa1111111111111
 end 
 
 else if @p_tblname = 'ITF_CUACCT_FUND'
 begin
    delete from dbo.ITF_CUACCT_FUND  where SYNC_TYPE = '1' and vipa1111111111111 = @p_vipa1111111111111 

	select custid, fundid, isnull(sum(payqty * matchprice),0.00) as fundeffect 
		into #fundeffect
		from dbo.mateconfirm 
		where issuetype = '1' and a333333333333 = '1'
		group by custid, fundid

    insert into dbo.ITF_CUACCT_FUND
       (USER_CODE, USER_NAME, CUACCT_CODE, CUACCT_ATTR, CUACCT_CLS, 
        CUACCT_LVL, CUACCT_GRP, CUACCT_DMF, INT_ORG, CURRENCY, 
        FUND_PREBLN, FUND_BLN, FUND_AVL, FUND_FRZ, FUND_UFZ, 
        FUND_TRD_FRZ, FUND_TRD_UFZ, FUND_TRD_OTD, FUND_TRD_BLN, FUND_CLR_FRZ, 
        FUND_CLR_UFZ, FUND_CLR_OTD, FUND_CLR_CTF, FUND_TRN_FRZ, FUND_TRN_UFZ, 
        FUND_DEBT, FUND_CREDIT, INT_RATE_SN, DR_RATE_SN, INT_CAL_BLN, 
        INT_BLN_ACCU, INT_CAL_DATE, INTEREST, INT_TAX, DR_INT, 
        MKT_VAL, CASH_ACCU, CHEQUE_ACCU, LAST_FUND_CLR, LWLMT_FUND, 
        LWLMT_ASSET, UPD_TIME, FUND_a333333333333, MAC, PAYLATER, 
        SUBSYS_SN, vipa1111111111111, SYNC_TYPE, CREDIT_FUND_BLN, FUND_SPECIAL,ASSETCASHPRO)
     select
        a.custid	                                as USER_CODE,
        ISNULL(ltrim(rtrim(b.fundname)),' ')        as USER_NAME,
        a.fundid                                    as CUACCT_CODE,
        '0'                                         as CUACCT_ATTR,
        case when len(b.fundkind) < 1 then ' ' else b.fundkind end     as CUACCT_CLS,
        case when len(b.fundlevel) < 1 then ' ' else b.fundlevel end   as CUACCT_LVL,
        case when len(b.fundgroup) < 1 then ' ' else b.fundgroup end   as CUACCT_GRP,
        '1'                                         as CUACCT_DMF,
        b.orgid                                     as INT_ORG,
        a.moneytype                                 as CURRENCY,
        a.fundlastbal                               as FUND_PREBLN,
        a.fundbal                                   as FUND_BLN,
        a.fundavl                                   as FUND_AVL,
        a.fundfrz                                   as FUND_FRZ,
        a.fundunfrz                                 as FUND_UFZ,
        a.fundbuy                                   as FUND_TRD_FRZ,
        a.fundsale                                  as FUND_TRD_UFZ,
        0                                           as FUND_TRD_OTD,
        a.fundbuysale                               as FUND_TRD_BLN,
        a.funduncomebuy                             as FUND_CLR_FRZ,
        a.funduncomesale                            as FUND_CLR_UFZ,
        0                                           as FUND_CLR_OTD,
        0                                           as FUND_CLR_CTF,
        0                                           as FUND_TRN_FRZ,
        0                                           as FUND_TRN_UFZ,
        0                                           as FUND_DEBT,
        0                                           as FUND_CREDIT,
        0                                           as INT_RATE_SN,
        0                                           as DR_RATE_SN,
        0                                           as INT_CAL_BLN,
        0                                           as INT_BLN_ACCU,
        0                                           as INT_CAL_DATE,
        0                                           as INTEREST,
        0                                           as INT_TAX,
        0                                           as DR_INT,
        a.marketvalue                               as MKT_VAL,
        0                                           as CASH_ACCU,
        0                                           as CHEQUE_ACCU,
        0                                           as LAST_FUND_CLR,
        0                                           as LWLMT_FUND,
        0                                           as LWLMT_ASSET,
        ' '                                         as UPD_TIME,
        '0'                                         as FUND_a333333333333,
        ' '                                         as MAC,
        0                                           as PAYLATER,
        @a444444444444444                                   as SUBSYS_SN,
        @p_vipa1111111111111                              as vipa1111111111111,
        '1'                                         as SYNC_TYPE,
		0										    as CREDIT_FUND_BLN,
		isnull(c.fundeffect,0.00)					as FUND_SPECIAL,
		a.fundcashpro								as assetcashpro
  from dbo.syncfundasset a 
   inner join dbo.vip_customers d with (nolock) on a.custid = d.custid and a.fundid = d.fundid and a.orgid = d.orgid and a.a1111111111111 = d.a1111111111111
   inner join dbo.fundinfo b with (nolock) on a.fundid = b.fundid and a.a1111111111111 = b.a1111111111111 and a.orgid = b.orgid and a.custid = b.custid
   left join #fundeffect c on b.fundid = c.fundid and b.custid = c.fundid
  where 1 = 1 and b.a333333333333 != '*'
    and d.vipid = @vipid
    and d.a333333333333 = '3'
    if @@error != 0
    begin
       select errorcode = -1, errormsg = '失败'
       return -1
    end
    
    select errorcode = 0 , errormsg = '成功'
    update dbo.itf_sync_tables set prepcnt = (select count(*) from dbo.ITF_CUACCT_FUND where SYNC_TYPE = '1' and vipa1111111111111 = @p_vipa1111111111111),a333333333333 = '1'
                where vipid = @vipid and tablename = @p_tblname and direct = 'e'
    select tablecnt = count(*) from dbo.ITF_CUACCT_FUND where SYNC_TYPE = '1' and vipa1111111111111 = @p_vipa1111111111111
 end 
 
 else if @p_tblname = 'ITF_STK_ASSET'
 begin
    delete from dbo.ITF_STK_ASSET  where SYNC_TYPE = '1' and vipa1111111111111 = @p_vipa1111111111111 
    insert into dbo.ITF_STK_ASSET
       (CUST_CODE, CUACCT_CODE, CUACCT_ATTR, CUACCT_CLS, CUACCT_LVL, 
        CUACCT_GRP, INT_ORG, STKEX, STKBD, STKPBU, 
        TRDACCT, CURRENCY, STK_CODE, STK_CLS, STK_PREBLN, 
        STK_BLN, STK_AVL, STK_FRZ, STK_UFZ, STK_NTRD, 
        STK_TRD_FRZ, STK_TRD_UFZ, STK_TRD_OTD, STK_TRD_BLN, STK_CLR_FRZ, 
        STK_CLR_UFZ, STK_CLR_OTD, STK_BCOST, STK_BCOST_RLT, STK_PLAMT, 
        STK_PLAMT_RLT, MKT_VAL, UPD_TIME, MAC, SUBSYS_SN, 
        vipa1111111111111, SYNC_TYPE)
    select
        a.custid                                                                                            as CUST_CODE,
        a.fundid                                                                                            as CUACCT_CODE,
        '0'                                                                                                 as CUACCT_ATTR,
        (select fundkind from dbo.fundinfo where a.fundid=fundid and fundtype != 'T' and a333333333333 != '*')    as CUACCT_CLS,
        (select fundlevel from dbo.fundinfo where a.fundid=fundid and fundtype != 'T' and a333333333333 != '*')   as CUACCT_LVL,
        (select fundgroup from dbo.fundinfo where a.fundid=fundid and fundtype != 'T' and a333333333333 != '*')   as CUACCT_GRP,
        a.orgid                                                                                             as INT_ORG,
        dbo.fn_ksts_market2STKEX(a.market)                                                              as STKEX,
        dbo.fn_ksts_market2STKBD(a.market, '0')                                                                    as STKBD,
        case when f.seat = '' then i.seat else f.seat end                                                   as STKPBU,        
        a.secuid                                                                                            as TRDACCT,
        a.moneytype                                                                                         as CURRENCY,
        a.stkcode                                                                                           as STK_CODE,
        dbo.fn_ksts_stktype2STK_CLS(g.stktype, '0', 0,g.market, ' ', g.isreg)                               as STK_CLS,
        a.stklastbal                                                                                        as STK_PREBLN,
        a.stkbal                                                                                            as STK_BLN,
        a.stkavl                                                                                            as STK_VAL,
        a.stkfrz                                                                                            as STK_FRZ,
        a.stkunfrz                                                                                          as STK_UFZ,
        0                                                                                                   as STK_NTRD,
        a.stkbuy                                                                                            as STK_TRD_FRZ,
        a.stksale	                                                                                        as STK_TRD_UFZ,
        --a.stkuncomebuy + a.stkuncomesale                                                                    as STK_TRD_OTD,
        0                                                                                                   as STK_TRD_OTD,
        a.stkbuysale                                                                                        as STK_TRD_BLN,
        a.stkuncomesale                                                                                     as STK_CLR_FRZ,
        a.stkuncomebuy                                                                                      as STK_CLR_UFZ,
        a.stkuncomebuy                                                                                      as STK_CLR_OTD,
        a.buycost                                                                                           as STK_BCOST,
        a.buycost                                                                                           as STK_BCOST_RLT,
        a.profitcost                                                                                        as STK_PLAMT,
        a.profitcost                                                                                        as STK_PLAMT_RLT,
        a.mktval                                                                                            as MKT_VAL,
        ' '                                                                                                 as UPD_TIME,
        ' '                                                                                                 as MAC,
        @a444444444444444                                                                                           as SUBSYS_SN,
        @p_vipa1111111111111                                                                                      as vipa1111111111111,
        '1'                                                                                                 as SYNC_TYPE
    FROM dbo.syncstkasset a  
    inner join dbo.fundinfo b with (nolock) on a.a1111111111111 = b.a1111111111111 and a.orgid = b.orgid and a.custid = b.custid and a.fundid = b.fundid and b.fundtype != 'T'
    inner join dbo.vip_customers c with (nolock) on b.a1111111111111 = c.a1111111111111 and b.orgid = c.orgid and b.custid = c.custid and b.fundid = c.fundid
    inner join dbo.secuid f with (nolock) on a.a1111111111111 = f.a1111111111111 and a.orgid = f.orgid and a.custid = f.custid and a.market = f.market and a.secuid = f.secuid
    inner join dbo.stktrd g with (nolock) on a.market = g.market and a.stkcode = g.stkcode and a.a1111111111111 = g.a1111111111111
    inner join dbo.orgseat i    on a.orgid = i.orgid and a.market = i.market and (i.pt_bsflag = '@' or charindex('0B', i.pt_bsflag) > 0 )
    where 1 = 1 and b.a333333333333 != '*' and f.a333333333333 != '*'
      and c.vipid = @vipid
      and c.a333333333333 = '3'
    if @@error != 0
    begin
       select errorcode = -1, errormsg = '失败'
       return -1
    end
    

    select errorcode = 0 , errormsg = 'ITF_STK_ASSET数据处理成功'
    update dbo.itf_sync_tables set prepcnt = (select count(*) from dbo.ITF_STK_ASSET where SYNC_TYPE = '1' and vipa1111111111111 = @p_vipa1111111111111),a333333333333 = '1'
           where vipid = @vipid and tablename = @p_tblname and direct = 'e'
    select tablecnt = count(*) from dbo.ITF_STK_ASSET where SYNC_TYPE = '1' and vipa1111111111111 = @p_vipa1111111111111
 end 
 
 else if @p_tblname = 'ITF_STK_COM_SPCFG'
 begin
    delete from dbo.ITF_STK_COM_SPCFG  where SYNC_TYPE = '1' and vipa1111111111111 = @p_vipa1111111111111 
    
     insert #stktype(a1111111111111,market,stktype, feestktype,moneytype)
    select distinct a1111111111111,market,stktype,feestktype = case when feestktype  = '' then stktype else feestktype end,moneytype 
       from dbo.stktype where stktype in (select stktype from dbo.stktype where feestktype != '' and feestktype != stktype)
        and a1111111111111 = @a1111111111111
        
     insert #diffstktype
             ( market, stktype )
    select distinct market,stktype from #stktype
        
    insert #stktype (a1111111111111,market,stktype,feestktype,moneytype) 
    select distinct a1111111111111,market,stktype,stktype,moneytype
       from dbo.stktype a
       where exists  (select 1 from #stktype where market = a.market and feestktype = a.stktype and a1111111111111 = a.a1111111111111 and moneytype = a.moneytype and feestktype <> stktype)

    select a.custid,a.fundid, a.orgid into #effectcust  from dbo.vip_customers a, dbo.customer b, dbo.fundinfo c
    where a.custid = b.custid and a.vipid = @vipid
     and a.custid = c.custid and a.fundid = c.fundid and a.orgid = c.orgid 
     and  c.fundtype != 'T' and c.a333333333333 != '*' and a.a333333333333 = '3'
      and (exists(select 1 from dbo.customer where custid = a.custid and calfeetype = '1' and effectdate <= @sysdate )
       or exists(select 1 from dbo.kfee_org_calflag where a.orgid = orgid and startdate <= @sysdate and kfeetype = '1'))
   
    create index #effectcust_idx_custid on #effectcust(custid)  

    select A.* into #surcharge_nocontain from 
    (
        select b.a1111111111111,a.orgid, a.custid,a.fundid, d.market, d.stktype, d.trdid, d.bsflag, d.stkcode, s.feerate, s.feemax, s.feemin, rowid = row_number()
            over(partition by b.a1111111111111, a.custid, d.market, d.stktype, d.trdid, d.bsflag, d.stkcode order by feerate desc)
        from #effectcust a, dbo.kfee_custcommission b, dbo.kfee_productdetails c, dbo.kfee_business d, dbo.kfee_custsurcharge e, dbo.kfee_surcharge s
        where a.custid = b.custid 
            and b.productid = c.productid and b.a1111111111111 = c.a1111111111111 
            and c.businessid = d.businessid and c.a1111111111111 = d.a1111111111111
            and b.custid = e.custid and b.productid = e.productid and b.a1111111111111 = e.a1111111111111 and charindex(e.servicetype, b.servicetypes) > 0
            and s.surchargeid = e.surchargeid and s.a1111111111111 = e.a1111111111111 and s.containflag = '0' 
    ) A where A.rowid = 1
    
    insert #bizConvert values(249,249) 
    insert #bizConvert values(249,251)
    insert #bizConvert values(249,253)
    insert #bizConvert values(250,250)
    insert #bizConvert values(250,252)
    insert #bizConvert values(250,254)
    
    insert #bizConvert values(242,960)
    insert #bizConvert values(243,961)
    
    insert #bizConvert values(100,100) 
    insert #bizConvert values(101,101)
    insert #bizConvert values(100,240)
    insert #bizConvert values(101,241)
    
    insert #bizConvert values(242,242)
    insert #bizConvert values(242,244)
    insert #bizConvert values(242,247)
    insert #bizConvert values(242,257)
    insert #bizConvert values(243,243)
    insert #bizConvert values(243,245)
    insert #bizConvert values(243,248)
    insert #bizConvert values(243,255)
    insert #bizConvert values(243,256)
    
	select rec_dentity=identity(int,1,1),finalstkBiz = isnull(b.tstkbiz, t1.STK_BIZ),t1.* into #ITF_STK_COM_SPCFG1 from 
    (
        select 
            REC_SN       = 10000000,
            INT_ORG      = convert(int, A.orgid),
            OBJ_TYPE     = '0',
            OBJ_VALUE    = convert(varchar(20), A.fundid),
            STKBD        = dbo.fn_ksts_market2STKBD(A.market, '0'),
            STK_CLS      = dbo.fn_ksts_stktype2STK_CLS(isnull(B.stktype,A.stktype), '0', 0, A.market,' ', ' '),
            STK_FLAG     = '0',
            STK_CODE     = '@',
            STK_BIZ      = dbo.fn_ksts_bsflag2STK_BIZ(isnull(B.stktype,A.stktype), A.trdid, A.bsflag, A.market),
            CHANNELS     = A.pt_operway,
            FEE_RATE     = A.commissionrate ,
            FEE_MAX      = A.maxfee,
            FEE_MIN      = A.minfee,
            vipa1111111111111  = @p_vipa1111111111111,
            SYNC_TYPE    =  '1'
        from dbo.specialcommission A		
        inner join dbo.vip_customers C with (nolock) on A.a1111111111111 = C.a1111111111111 and A.orgid = C.orgid and A.fundid = C.fundid and A.custid = C.custid
        inner join dbo.fundinfo F with (nolock) on C.a1111111111111 = F.a1111111111111 and C.orgid = F.orgid and C.custid = F.custid and C.fundid = F.fundid
        inner join dbo.customer k on A.custid = k.custid and A.a1111111111111 = k.a1111111111111
        left join #stktype B  on A.market = B.market and A.moneytype = B.moneytype and A.stktype = B.feestktype
        where F.fundtype != 'T' and F.a333333333333 != '*' and C.vipid = @vipid
            and A.trdid not in('4', '7', '8', 'R') and C.a333333333333 = '3' 
            and ((A.market+A.stktype in (select market+stktype from #diffstktype) and B.stktype is not null) or A.market+A.stktype not in (select market+stktype from #diffstktype))

            and (not (k.calfeetype = '1' and k.effectdate <= @sysdate))
            and k.orgid not in (select orgid from dbo.kfee_org_calflag where startdate <= @sysdate and kfeetype = '1')
        union all 
        select           REC_SN       = 10000000,
            INT_ORG      = convert(int, A.orgid),
            OBJ_TYPE     = '0',
            OBJ_VALUE    = convert(varchar(20), A.fundid),
            STKBD        = dbo.fn_ksts_market2STKBD(A.market, '0'),
            STK_CLS      = dbo.fn_ksts_stktype2STK_CLS(isnull(B.stktype,A.stktype), '0', 0,A.market,' ', ' '),
            STK_FLAG     = '0',
            STK_CODE     = A.stkcode,
            STK_BIZ      = dbo.fn_ksts_bsflag2STK_BIZ(isnull(B.stktype,A.stktype), A.trdid, A.bsflag, A.market),
            CHANNELS     = '@',
            FEE_RATE     = A.feerate + isnull(s.feerate, 0.00), 
            FEE_MAX      = A.feemax,
            FEE_MIN      = A.feemin,
            vipa1111111111111  = @p_vipa1111111111111,
            SYNC_TYPE    =  '1'    
        from 
        (  
            select a.*, rowid = row_number() over(partition by a.a1111111111111,a.custid, a.market, a.stktype, a.trdid, a.bsflag,a.stkcode order by feerate desc) from 
            (
                select a.a1111111111111,c.orgid, a.custid,c.fundid, b.market,b.stktype, b.trdid, b.bsflag, b.stkcode, k.feerate, k.feemax, k.feemin
                from dbo.kfee_custbusiness a, dbo.kfee_business b,#effectcust c, dbo.kfee_productprices k  
                where a.businessid = b.businessid and a.a1111111111111 = b.a1111111111111 and b.stkcode = ''
                    and a.custid = c.custid 
                    and a.priceid = k.priceid and a.a1111111111111 = k.a1111111111111
                union
                select a.a1111111111111, d.orgid, a.custid,d.fundid, c.market,c.stktype, c.trdid, c.bsflag, c.stkcode, k.feerate, k.feemax, k.feemin
                from dbo.kfee_custcommission a, dbo.kfee_productdetails b, dbo.kfee_business c,#effectcust d, dbo.kfee_productprices k  
                where a.a1111111111111 = b.a1111111111111 and a.productid = b.productid and c.stkcode = ''
                    and b.businessid = c.businessid and b.a1111111111111 = c.a1111111111111
                    and a.custid = d.custid
                    and a.priceid = k.priceid and a.a1111111111111 = k.a1111111111111
                union
                select b.a1111111111111, a.orgid, a.custid,a.fundid, d.market, d.stktype, d.trdid, d.bsflag, d.stkcode, s.feerate, s.feemax, s.feemin
                from #effectcust a, dbo.kfee_custcommission b, dbo.kfee_productdetails c, dbo.kfee_business d, dbo.kfee_custsurcharge e, dbo.kfee_surcharge s
                where a.custid = b.custid 
                    and b.productid = c.productid and b.a1111111111111 = c.a1111111111111 
                    and c.businessid = d.businessid and c.a1111111111111 = d.a1111111111111
                    and b.custid = e.custid and b.productid = e.productid and b.a1111111111111 = e.a1111111111111 and charindex(e.servicetype, b.servicetypes) > 0
                    and s.surchargeid = e.surchargeid and s.a1111111111111 = e.a1111111111111 and s.containflag = '1' 
            ) a 
        ) A  inner join dbo.marketmoney e on A.market = e.market and A.a1111111111111 = e.a1111111111111
        left join #stktype B  on A.market = B.market and e.moneytype = B.moneytype and A.stktype = B.feestktype
        left join #surcharge_nocontain s on s.a1111111111111 = A.a1111111111111 and s.custid = A.custid and s.market = A.market and s.stktype = A.stktype and s.trdid = A.trdid and s.bsflag = A.bsflag and s.stkcode = A.stkcode
        where A.a1111111111111 = @a1111111111111 and A.rowid = 1
            and A.trdid not in('4', '7', '8', 'R')
            and ((A.market+A.stktype in (select market+stktype from #diffstktype) and B.stktype is not null) or A.market+A.stktype not in (select market+stktype from #diffstktype))

	) t1 left join #bizConvert b on STK_BIZ = b.fstkbiz order by OBJ_VALUE
     
    insert into dbo.ITF_STK_COM_SPCFG
       (REC_SN, INT_ORG, OBJ_TYPE, OBJ_VALUE, STKBD, 
        STK_CLS, STK_FLAG, STK_CODE, STK_BIZ, CHANNELS, 
        FEE_RATE, FEE_MAX, FEE_MIN, vipa1111111111111, SYNC_TYPE)
    	select rec_dentity+REC_SN, INT_ORG, OBJ_TYPE, OBJ_VALUE, STKBD, 
        STK_CLS, STK_FLAG, STK_CODE, finalstkBiz, CHANNELS, 
        FEE_RATE, FEE_MAX, FEE_MIN, vipa1111111111111, SYNC_TYPE
      from #ITF_STK_COM_SPCFG1 
      
    if @@error != 0
    begin
       select errorcode = -1, errormsg = '失败'
       return -1
    end
    
    drop table #ITF_STK_COM_SPCFG1
    
    
    if exists (select * from dbo.ITF_STK_COM_SPCFG where STK_BIZ in (181, 182, 183) and SYNC_TYPE = '1' and vipa1111111111111 = @p_vipa1111111111111)
    begin
    	 
      select REC_SN=identity(int,1,1),t2.* into #ITF_STK_COM_SPCFG2 from (  
	    	select
	           INT_ORG      = A.INT_ORG,
	           OBJ_TYPE     = A.OBJ_TYPE,
	           OBJ_VALUE    = A.OBJ_VALUE,
	           STKBD        = A.STKBD,
	           STK_CLS      = A.STK_CLS,
	           STK_FLAG     = A.STK_FLAG,
	           STK_CODE     = A.STK_CODE,
	           STK_BIZ      = case when A.STK_BIZ = 181 then 187 when A.STK_BIZ = 182 then 188 else 180 end,
	           CHANNELS     = A.CHANNELS,
	           FEE_RATE     = A.FEE_RATE,
	           FEE_MAX      = A.FEE_MAX,
	           FEE_MIN      = A.FEE_MIN,
	           vipa1111111111111  = @p_vipa1111111111111,
	           SYNC_TYPE    =  '1'
    		from dbo.ITF_STK_COM_SPCFG A where A.STK_BIZ in (181, 182, 183) and A.SYNC_TYPE = '1' and A.vipa1111111111111 = @p_vipa1111111111111) t2 order by OBJ_VALUE
    	
    	insert into dbo.ITF_STK_COM_SPCFG
      		(REC_SN, INT_ORG, OBJ_TYPE, OBJ_VALUE, STKBD, 
       		 STK_CLS, STK_FLAG, STK_CODE, STK_BIZ, CHANNELS, 
        		FEE_RATE, FEE_MAX, FEE_MIN, vipa1111111111111, SYNC_TYPE)
     		select REC_SN + 30000000, INT_ORG, OBJ_TYPE, OBJ_VALUE, STKBD, 
        		STK_CLS, STK_FLAG, STK_CODE, STK_BIZ, CHANNELS, 
        		FEE_RATE, FEE_MAX, FEE_MIN, vipa1111111111111, SYNC_TYPE
        from  #ITF_STK_COM_SPCFG2   
      if @@error != 0
      begin
          select errorcode = -1, errormsg = '失败'
         return -1
      end
        
        drop table #ITF_STK_COM_SPCFG2
    end
    
    if @@error != 0
    begin
       select errorcode = -1, errormsg = '失败'
       return -1
    end
    
    if exists (select * from dbo.ITF_STK_COM_SPCFG where STK_BIZ in (181, 182, 183) and SYNC_TYPE = '1' and vipa1111111111111 = @p_vipa1111111111111)
    begin
    	 
      select REC_SN=identity(int,1,1),t3.* into #ITF_STK_COM_SPCFG3 from (
	    	select 
	           INT_ORG      = A.INT_ORG,
	           OBJ_TYPE     = A.OBJ_TYPE,
	           OBJ_VALUE    = A.OBJ_VALUE,
	           STKBD        = A.STKBD,
	           STK_CLS      = A.STK_CLS,
	           STK_FLAG     = A.STK_FLAG,
	           STK_CODE     = A.STK_CODE,
	           STK_BIZ      = case when A.STK_BIZ = 181 then 827 when A.STK_BIZ = 182 then 828 else 184 end,
	           CHANNELS     = A.CHANNELS,
	           FEE_RATE     = A.FEE_RATE,
	           FEE_MAX      = A.FEE_MAX,
	           FEE_MIN      = A.FEE_MIN,
	           vipa1111111111111  = @p_vipa1111111111111,
	           SYNC_TYPE    =  '1'
    	from dbo.ITF_STK_COM_SPCFG A where A.STK_BIZ in (181, 182, 183) and A.SYNC_TYPE = '1' and A.vipa1111111111111 = @p_vipa1111111111111) t3 order by OBJ_VALUE
    	
    	insert into dbo.ITF_STK_COM_SPCFG
       	(REC_SN, INT_ORG, OBJ_TYPE, OBJ_VALUE, STKBD, 
        	STK_CLS, STK_FLAG, STK_CODE, STK_BIZ, CHANNELS, 
        	FEE_RATE, FEE_MAX, FEE_MIN, vipa1111111111111, SYNC_TYPE)
      	select REC_SN + 35000000, INT_ORG, OBJ_TYPE, OBJ_VALUE, STKBD, 
        	STK_CLS, STK_FLAG, STK_CODE, STK_BIZ, CHANNELS, 
        	FEE_RATE, FEE_MAX, FEE_MIN, vipa1111111111111, SYNC_TYPE
        	from #ITF_STK_COM_SPCFG3
        	
       if @@error != 0
      begin
          select errorcode = -1, errormsg = '失败'
         return -1
      end
        	
      drop table #ITF_STK_COM_SPCFG3
        
    end
    if exists (select * from dbo.ITF_STK_COM_SPCFG where STK_BIZ in (181, 182) and SYNC_TYPE = '1' and vipa1111111111111 = @p_vipa1111111111111)
    begin
      select REC_SN=identity(int,1,1),t3.* into #ITF_STK_COM_SPCFG825_826 from (
         select 
             INT_ORG      = A.INT_ORG,
             OBJ_TYPE     = A.OBJ_TYPE,
             OBJ_VALUE    = A.OBJ_VALUE,
             STKBD        = A.STKBD,
             STK_CLS      = A.STK_CLS,
             STK_FLAG     = A.STK_FLAG,
             STK_CODE     = A.STK_CODE,
             STK_BIZ      = case when A.STK_BIZ = 181 then 825 else 826 end,
             CHANNELS     = A.CHANNELS,
             FEE_RATE     = A.FEE_RATE,
             FEE_MAX      = A.FEE_MAX,
             FEE_MIN      = A.FEE_MIN,
             vipa1111111111111  = @p_vipa1111111111111,
             SYNC_TYPE    =  '1'
      from dbo.ITF_STK_COM_SPCFG A where A.STK_BIZ in (181, 182) and A.SYNC_TYPE = '1' and A.vipa1111111111111 = @p_vipa1111111111111) t3 order by OBJ_VALUE
      
      insert into dbo.ITF_STK_COM_SPCFG
        (REC_SN, INT_ORG, OBJ_TYPE, OBJ_VALUE, STKBD, 
           STK_CLS, STK_FLAG, STK_CODE, STK_BIZ, CHANNELS, 
           FEE_RATE, FEE_MAX, FEE_MIN, vipa1111111111111, SYNC_TYPE)
         select REC_SN + 40000000, INT_ORG, OBJ_TYPE, OBJ_VALUE, STKBD, 
          STK_CLS, STK_FLAG, STK_CODE, STK_BIZ, CHANNELS, 
          FEE_RATE, FEE_MAX, FEE_MIN, vipa1111111111111, SYNC_TYPE
          from #ITF_STK_COM_SPCFG825_826
      if @@error != 0
      begin
         select errorcode = -1, errormsg = '失败'
         return -1
      end
      drop table #ITF_STK_COM_SPCFG825_826
        
    end
    if exists (select * from dbo.ITF_STK_COM_SPCFG where STK_BIZ in (803, 804) and SYNC_TYPE = '1' and vipa1111111111111 = @p_vipa1111111111111)
    begin
      select REC_SN=identity(int,1,1),t6.* into #ITF_STK_COM_SPCFG807_808 from (
         select 
             INT_ORG      = A.INT_ORG,
             OBJ_TYPE     = A.OBJ_TYPE,
             OBJ_VALUE    = A.OBJ_VALUE,
             STKBD        = A.STKBD,
             STK_CLS      = A.STK_CLS,
             STK_FLAG     = A.STK_FLAG,
             STK_CODE     = A.STK_CODE,
             STK_BIZ      = case when A.STK_BIZ = 803 then 807 else 808 end,
             CHANNELS     = A.CHANNELS,
             FEE_RATE     = A.FEE_RATE,
             FEE_MAX      = A.FEE_MAX,
             FEE_MIN      = A.FEE_MIN,
             vipa1111111111111  = @p_vipa1111111111111,
             SYNC_TYPE    =  '1'
      from dbo.ITF_STK_COM_SPCFG A where A.STK_BIZ in (803, 804) and A.SYNC_TYPE = '1' and A.vipa1111111111111 = @p_vipa1111111111111) t6 order by OBJ_VALUE
      
      insert into dbo.ITF_STK_COM_SPCFG
        (REC_SN, INT_ORG, OBJ_TYPE, OBJ_VALUE, STKBD, 
           STK_CLS, STK_FLAG, STK_CODE, STK_BIZ, CHANNELS, 
           FEE_RATE, FEE_MAX, FEE_MIN, vipa1111111111111, SYNC_TYPE)
         select REC_SN + 45000000, INT_ORG, OBJ_TYPE, OBJ_VALUE, STKBD, 
          STK_CLS, STK_FLAG, STK_CODE, STK_BIZ, CHANNELS, 
          FEE_RATE, FEE_MAX, FEE_MIN, vipa1111111111111, SYNC_TYPE
          from #ITF_STK_COM_SPCFG807_808
      if @@error != 0
      begin
         select errorcode = -1, errormsg = '失败'
         return -1
      end
      drop table #ITF_STK_COM_SPCFG807_808
        
    end
	if exists (select * from dbo.ITF_STK_COM_SPCFG where STK_BIZ in (260,261,960,961,966,967)and SYNC_TYPE = '1' and vipa1111111111111 = @p_vipa1111111111111)
    begin
    	 
      select REC_SN=identity(int,1,1),t4.* into #ITF_STK_COM_SPCFG4 from (
	    	select 
	           INT_ORG      = A.INT_ORG,
	           OBJ_TYPE     = A.OBJ_TYPE,
	           OBJ_VALUE    = A.OBJ_VALUE,
	           STKBD        = A.STKBD,
	           STK_CLS      = A.STK_CLS,
	           STK_FLAG     = A.STK_FLAG,
	           STK_CODE     = A.STK_CODE,
	           STK_BIZ      =	case	when A.STK_BIZ = 260 then 262 
										when A.STK_BIZ = 261 then 263  
										when A.STK_BIZ = 960 then 962
										when A.STK_BIZ = 961 then 963
										when A.STK_BIZ = 966 then 968
										when A.STK_BIZ = 967 then 969 
								end,
	           CHANNELS     = A.CHANNELS,
	           FEE_RATE     = A.FEE_RATE,
	           FEE_MAX      = A.FEE_MAX,
	           FEE_MIN      = A.FEE_MIN,
	           vipa1111111111111  = A.vipa1111111111111,
	           SYNC_TYPE    = A.SYNC_TYPE
    	from dbo.ITF_STK_COM_SPCFG A where A.STK_BIZ in (260,261,960,961,966,967)and SYNC_TYPE = '1' and vipa1111111111111 = @p_vipa1111111111111) t4 order by OBJ_VALUE
    	
    	insert into dbo.ITF_STK_COM_SPCFG
       	(REC_SN, INT_ORG, OBJ_TYPE, OBJ_VALUE, STKBD, 
        	STK_CLS, STK_FLAG, STK_CODE, STK_BIZ, CHANNELS, 
        	FEE_RATE, FEE_MAX, FEE_MIN, vipa1111111111111, SYNC_TYPE)
      	select REC_SN + 50000000, INT_ORG, OBJ_TYPE, OBJ_VALUE, STKBD, 
        	STK_CLS, STK_FLAG, STK_CODE, STK_BIZ, CHANNELS, 
        	FEE_RATE, FEE_MAX, FEE_MIN, vipa1111111111111, SYNC_TYPE
        	from #ITF_STK_COM_SPCFG4
      if @@error != 0
      begin
         select errorcode = -1, errormsg = '失败1'
         return -1
      end 	
      drop table #ITF_STK_COM_SPCFG4
	  

    	 
      select REC_SN=identity(int,1,1),t5.* into #ITF_STK_COM_SPCFG5 from (
	    	select 
	           INT_ORG      = A.INT_ORG,
	           OBJ_TYPE     = A.OBJ_TYPE,
	           OBJ_VALUE    = A.OBJ_VALUE,
	           STKBD        = A.STKBD,
	           STK_CLS      = A.STK_CLS,
	           STK_FLAG     = A.STK_FLAG,
	           STK_CODE     = A.STK_CODE,
	           STK_BIZ      =	case	when A.STK_BIZ = 260 then 264
										when A.STK_BIZ = 261 then 265
										when A.STK_BIZ = 262 then 268
										when A.STK_BIZ = 263 then 269
										when A.STK_BIZ = 960 then 964
										when A.STK_BIZ = 961 then 965
										when A.STK_BIZ = 962 then 970
										when A.STK_BIZ = 963 then 971
								end,
	           CHANNELS     = A.CHANNELS,
	           FEE_RATE     = A.FEE_RATE,
	           FEE_MAX      = A.FEE_MAX,
	           FEE_MIN      = A.FEE_MIN,
	           vipa1111111111111  = A.vipa1111111111111,
	           SYNC_TYPE    = A.SYNC_TYPE
    	from dbo.ITF_STK_COM_SPCFG A where A.STK_BIZ in (260,261,262,263,960,961,962,963)and SYNC_TYPE = '1' and vipa1111111111111 = @p_vipa1111111111111) t5 order by OBJ_VALUE
    	
    	insert into dbo.ITF_STK_COM_SPCFG
       	(REC_SN, INT_ORG, OBJ_TYPE, OBJ_VALUE, STKBD, 
        	STK_CLS, STK_FLAG, STK_CODE, STK_BIZ, CHANNELS, 
        	FEE_RATE, FEE_MAX, FEE_MIN, vipa1111111111111, SYNC_TYPE)
      	select REC_SN + 55000000, INT_ORG, OBJ_TYPE, OBJ_VALUE, STKBD, 
        	STK_CLS, STK_FLAG, STK_CODE, STK_BIZ, CHANNELS, 
        	FEE_RATE, FEE_MAX, FEE_MIN, vipa1111111111111, SYNC_TYPE
        	from #ITF_STK_COM_SPCFG5
      if @@error != 0
      begin
         select errorcode = -1, errormsg = '失败2'
         return -1
      end  	
      drop table #ITF_STK_COM_SPCFG5
        
    end

    
    if exists (select * from dbo.ITF_STK_COM_SPCFG where STK_BIZ in (100,101) and SYNC_TYPE = '1' and vipa1111111111111 = @p_vipa1111111111111)
    begin
      select REC_SN=identity(int,1,1),t9.* into #ITF_STK_COM_SPCFG120_121 from (
         select 
             INT_ORG      = A.INT_ORG,
             OBJ_TYPE     = A.OBJ_TYPE,
             OBJ_VALUE    = A.OBJ_VALUE,
             STKBD        = A.STKBD,
             STK_CLS      = A.STK_CLS,
             STK_FLAG     = A.STK_FLAG,
             STK_CODE     = A.STK_CODE,
             STK_BIZ      = case when A.STK_BIZ = 100 then 120 else 121 end,
             CHANNELS     = A.CHANNELS,
             FEE_RATE     = A.FEE_RATE,
             FEE_MAX      = A.FEE_MAX,
             FEE_MIN      = A.FEE_MIN,
             vipa1111111111111  = @p_vipa1111111111111,
             SYNC_TYPE    =  '1'
      from dbo.ITF_STK_COM_SPCFG A where A.STK_BIZ in (100,101) and A.SYNC_TYPE = '1' and A.vipa1111111111111 = @p_vipa1111111111111) t9 order by OBJ_VALUE
      
      insert into dbo.ITF_STK_COM_SPCFG
        (REC_SN, INT_ORG, OBJ_TYPE, OBJ_VALUE, STKBD, 
           STK_CLS, STK_FLAG, STK_CODE, STK_BIZ, CHANNELS, 
           FEE_RATE, FEE_MAX, FEE_MIN, vipa1111111111111, SYNC_TYPE)
         select REC_SN + 60000000, INT_ORG, OBJ_TYPE, OBJ_VALUE, STKBD, 
          STK_CLS, STK_FLAG, STK_CODE, STK_BIZ, CHANNELS, 
          FEE_RATE, FEE_MAX, FEE_MIN, vipa1111111111111, SYNC_TYPE
          from #ITF_STK_COM_SPCFG120_121
      if @@error != 0
      begin
         select errorcode = -1, errormsg = '失败'
         return -1
      end
      drop table #ITF_STK_COM_SPCFG120_121
    end

 
    
    if @@error != 0
    begin
       select errorcode = -1 , errormsg = '失败'
    end
    
    select errorcode = 0 , errormsg = '成功'    
	update dbo.itf_sync_tables set prepcnt = (select count(*) from dbo.STK_HS_DETAIL where SYNC_TYPE = '1' and vipa1111111111111 = @p_vipa1111111111111),a333333333333 = '1'
                where vipid = @vipid and tablename = @p_tblname and direct = 'e'
    
    select tablecnt = count(*) from dbo.STK_HS_DETAIL where SYNC_TYPE = '1' and vipa1111111111111 = @p_vipa1111111111111
 end

 else
 begin
    select errorcode = -1, errormsg = '错误'
    return -1
 end
 
 return 0
go

exec nb_add_sysobjects 'P','up_com_test'
go