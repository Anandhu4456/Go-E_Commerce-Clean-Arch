//
// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// This is a commercial product and requires a license to operate.
// A trial license can be obtained at https://unidoc.io
//
// DO NOT EDIT: generated by unitwist Go source code obfuscator.
//
// Use of this source code is governed by the UniDoc End User License Agreement
// terms that can be accessed at https://unidoc.io/eula/

package arithmetic ;import (_g "fmt";_b "github.com/unidoc/unipdf/v3/common";_cg "github.com/unidoc/unipdf/v3/internal/bitwise";_d "github.com/unidoc/unipdf/v3/internal/jbig2/internal";_c "io";_e "strings";);func (_fed *DecoderStats )String ()string {_gaa :=&_e .Builder {};
_gaa .WriteString (_g .Sprintf ("S\u0074\u0061\u0074\u0073\u003a\u0020\u0020\u0025\u0064\u000a",len (_fed ._adb )));for _bcf ,_cce :=range _fed ._adb {if _cce !=0{_gaa .WriteString (_g .Sprintf ("N\u006f\u0074\u0020\u007aer\u006f \u0061\u0074\u003a\u0020\u0025d\u0020\u002d\u0020\u0025\u0064\u000a",_bcf ,_cce ));
};};return _gaa .String ();};func (_cc *Decoder )lpsExchange (_gaf *DecoderStats ,_cgdg int32 ,_ffe uint32 )int {_egd :=_gaf .getMps ();if _cc ._a < _ffe {_gaf .setEntry (int (_ce [_cgdg ][1]));_cc ._a =_ffe ;return int (_egd );};if _ce [_cgdg ][3]==1{_gaf .toggleMps ();
};_gaf .setEntry (int (_ce [_cgdg ][2]));_cc ._a =_ffe ;return int (1-_egd );};func (_cd *DecoderStats )cx ()byte {return _cd ._adb [_cd ._ed ]};func (_fg *Decoder )DecodeIAID (codeLen uint64 ,stats *DecoderStats )(int64 ,error ){_fg ._gc =1;var _bdc uint64 ;
for _bdc =0;_bdc < codeLen ;_bdc ++{stats .SetIndex (int32 (_fg ._gc ));_fc ,_gg :=_fg .DecodeBit (stats );if _gg !=nil {return 0,_gg ;};_fg ._gc =(_fg ._gc <<1)|int64 (_fc );};_gag :=_fg ._gc -(1<<codeLen );return _gag ,nil ;};func (_gce *Decoder )DecodeInt (stats *DecoderStats )(int32 ,error ){var (_dbe ,_bd int32 ;
_ff ,_bc ,_cgd int ;_ad error ;);if stats ==nil {stats =NewStats (512,1);};_gce ._gc =1;_bc ,_ad =_gce .decodeIntBit (stats );if _ad !=nil {return 0,_ad ;};_ff ,_ad =_gce .decodeIntBit (stats );if _ad !=nil {return 0,_ad ;};if _ff ==1{_ff ,_ad =_gce .decodeIntBit (stats );
if _ad !=nil {return 0,_ad ;};if _ff ==1{_ff ,_ad =_gce .decodeIntBit (stats );if _ad !=nil {return 0,_ad ;};if _ff ==1{_ff ,_ad =_gce .decodeIntBit (stats );if _ad !=nil {return 0,_ad ;};if _ff ==1{_ff ,_ad =_gce .decodeIntBit (stats );if _ad !=nil {return 0,_ad ;
};if _ff ==1{_cgd =32;_bd =4436;}else {_cgd =12;_bd =340;};}else {_cgd =8;_bd =84;};}else {_cgd =6;_bd =20;};}else {_cgd =4;_bd =4;};}else {_cgd =2;_bd =0;};for _ee :=0;_ee < _cgd ;_ee ++{_ff ,_ad =_gce .decodeIntBit (stats );if _ad !=nil {return 0,_ad ;
};_dbe =(_dbe <<1)|int32 (_ff );};_dbe +=_bd ;if _bc ==0{return _dbe ,nil ;}else if _bc ==1&&_dbe > 0{return -_dbe ,nil ;};return 0,_d .ErrOOB ;};type Decoder struct{ContextSize []uint32 ;ReferedToContextSize []uint32 ;_ga *_cg .Reader ;_dc uint8 ;_cb uint64 ;
_a uint32 ;_gc int64 ;_af int32 ;_ge int32 ;_ag int64 ;};func (_ded *Decoder )init ()error {_ded ._ag =_ded ._ga .AbsolutePosition ();_cgdc ,_dcb :=_ded ._ga .ReadByte ();if _dcb !=nil {_b .Log .Debug ("B\u0075\u0066\u0066\u0065\u0072\u0030 \u0072\u0065\u0061\u0064\u0042\u0079\u0074\u0065\u0020f\u0061\u0069\u006ce\u0064.\u0020\u0025\u0076",_dcb );
return _dcb ;};_ded ._dc =_cgdc ;_ded ._cb =uint64 (_cgdc )<<16;if _dcb =_ded .readByte ();_dcb !=nil {return _dcb ;};_ded ._cb <<=7;_ded ._af -=7;_ded ._a =0x8000;_ded ._ge ++;return nil ;};func (_def *DecoderStats )Overwrite (dNew *DecoderStats ){for _gef :=0;
_gef < len (_def ._adb );_gef ++{_def ._adb [_gef ]=dNew ._adb [_gef ];_def ._ca [_gef ]=dNew ._ca [_gef ];};};func New (r *_cg .Reader )(*Decoder ,error ){_ac :=&Decoder {_ga :r ,ContextSize :[]uint32 {16,13,10,10},ReferedToContextSize :[]uint32 {13,10}};
if _ab :=_ac .init ();_ab !=nil {return nil ,_ab ;};return _ac ,nil ;};func (_cgce *DecoderStats )toggleMps (){_cgce ._ca [_cgce ._ed ]^=1};func (_ef *Decoder )DecodeBit (stats *DecoderStats )(int ,error ){var (_cgc int ;_eg =_ce [stats .cx ()][0];_bg =int32 (stats .cx ());
);defer func (){_ef ._ge ++}();_ef ._a -=_eg ;if (_ef ._cb >>16)< uint64 (_eg ){_cgc =_ef .lpsExchange (stats ,_bg ,_eg );if _de :=_ef .renormalize ();_de !=nil {return 0,_de ;};}else {_ef ._cb -=uint64 (_eg )<<16;if (_ef ._a &0x8000)==0{_cgc =_ef .mpsExchange (stats ,_bg );
if _db :=_ef .renormalize ();_db !=nil {return 0,_db ;};}else {_cgc =int (stats .getMps ());};};return _cgc ,nil ;};func (_bf *Decoder )readByte ()error {if _bf ._ga .AbsolutePosition ()> _bf ._ag {if _ ,_bcg :=_bf ._ga .Seek (-1,_c .SeekCurrent );_bcg !=nil {return _bcg ;
};};_fff ,_dbf :=_bf ._ga .ReadByte ();if _dbf !=nil {return _dbf ;};_bf ._dc =_fff ;if _bf ._dc ==0xFF{_abb ,_cf :=_bf ._ga .ReadByte ();if _cf !=nil {return _cf ;};if _abb > 0x8F{_bf ._cb +=0xFF00;_bf ._af =8;if _ ,_dd :=_bf ._ga .Seek (-2,_c .SeekCurrent );
_dd !=nil {return _dd ;};}else {_bf ._cb +=uint64 (_abb )<<9;_bf ._af =7;};}else {_fff ,_dbf =_bf ._ga .ReadByte ();if _dbf !=nil {return _dbf ;};_bf ._dc =_fff ;_bf ._cb +=uint64 (_bf ._dc )<<8;_bf ._af =8;};_bf ._cb &=0xFFFFFFFFFF;return nil ;};var (_ce =[][4]uint32 {{0x5601,1,1,1},{0x3401,2,6,0},{0x1801,3,9,0},{0x0AC1,4,12,0},{0x0521,5,29,0},{0x0221,38,33,0},{0x5601,7,6,1},{0x5401,8,14,0},{0x4801,9,14,0},{0x3801,10,14,0},{0x3001,11,17,0},{0x2401,12,18,0},{0x1C01,13,20,0},{0x1601,29,21,0},{0x5601,15,14,1},{0x5401,16,14,0},{0x5101,17,15,0},{0x4801,18,16,0},{0x3801,19,17,0},{0x3401,20,18,0},{0x3001,21,19,0},{0x2801,22,19,0},{0x2401,23,20,0},{0x2201,24,21,0},{0x1C01,25,22,0},{0x1801,26,23,0},{0x1601,27,24,0},{0x1401,28,25,0},{0x1201,29,26,0},{0x1101,30,27,0},{0x0AC1,31,28,0},{0x09C1,32,29,0},{0x08A1,33,30,0},{0x0521,34,31,0},{0x0441,35,32,0},{0x02A1,36,33,0},{0x0221,37,34,0},{0x0141,38,35,0},{0x0111,39,36,0},{0x0085,40,37,0},{0x0049,41,38,0},{0x0025,42,39,0},{0x0015,43,40,0},{0x0009,44,41,0},{0x0005,45,42,0},{0x0001,45,43,0},{0x5601,46,46,0}};
);func (_gf *Decoder )renormalize ()error {for {if _gf ._af ==0{if _fa :=_gf .readByte ();_fa !=nil {return _fa ;};};_gf ._a <<=1;_gf ._cb <<=1;_gf ._af --;if (_gf ._a &0x8000)!=0{break ;};};_gf ._cb &=0xffffffff;return nil ;};func (_ced *DecoderStats )setEntry (_ec int ){_eaa :=byte (_ec &0x7f);
_ced ._adb [_ced ._ed ]=_eaa };type DecoderStats struct{_ed int32 ;_cbg int32 ;_adb []byte ;_ca []byte ;};func (_ddg *DecoderStats )SetIndex (index int32 ){_ddg ._ed =index };func (_aed *DecoderStats )Reset (){for _dga :=0;_dga < len (_aed ._adb );_dga ++{_aed ._adb [_dga ]=0;
_aed ._ca [_dga ]=0;};};func (_df *DecoderStats )Copy ()*DecoderStats {_ea :=&DecoderStats {_cbg :_df ._cbg ,_adb :make ([]byte ,_df ._cbg )};copy (_ea ._adb ,_df ._adb );return _ea ;};func (_cbc *Decoder )mpsExchange (_cfe *DecoderStats ,_dg int32 )int {_ae :=_cfe ._ca [_cfe ._ed ];
if _cbc ._a < _ce [_dg ][0]{if _ce [_dg ][3]==1{_cfe .toggleMps ();};_cfe .setEntry (int (_ce [_dg ][2]));return int (1-_ae );};_cfe .setEntry (int (_ce [_dg ][1]));return int (_ae );};func NewStats (contextSize int32 ,index int32 )*DecoderStats {return &DecoderStats {_ed :index ,_cbg :contextSize ,_adb :make ([]byte ,contextSize ),_ca :make ([]byte ,contextSize )};
};func (_aa *Decoder )decodeIntBit (_eea *DecoderStats )(int ,error ){_eea .SetIndex (int32 (_aa ._gc ));_fe ,_ba :=_aa .DecodeBit (_eea );if _ba !=nil {_b .Log .Debug ("\u0041\u0072\u0069\u0074\u0068\u006d\u0065t\u0069\u0063\u0044e\u0063\u006f\u0064e\u0072\u0020'\u0064\u0065\u0063\u006f\u0064\u0065I\u006etB\u0069\u0074\u0027\u002d\u003e\u0020\u0044\u0065\u0063\u006f\u0064\u0065\u0042\u0069\u0074\u0020\u0066\u0061\u0069\u006c\u0065\u0064\u002e\u0020\u0025\u0076",_ba );
return _fe ,_ba ;};if _aa ._gc < 256{_aa ._gc =((_aa ._gc <<uint64 (1))|int64 (_fe ))&0x1ff;}else {_aa ._gc =(((_aa ._gc <<uint64 (1)|int64 (_fe ))&511)|256)&0x1ff;};return _fe ,nil ;};func (_ede *DecoderStats )getMps ()byte {return _ede ._ca [_ede ._ed ]};