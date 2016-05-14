#include <stdio.h>
#include <string.h>
#include <unistd.h>
#include <stdlib.h>
#include <hamlib/rig.h>

hamlib_port_t myport;
RIG *myrig;

int set_port(int rig_port_type, char* portname, int baudrate, int databits, int stopbits, int parity, int handshake){

	//check if rig exists
 	if (myrig == 0) return -1;

	myrig->state.rigport.type.rig = rig_port_type;
	myrig->state.rigport.parm.serial.rate = baudrate;
	myrig->state.rigport.parm.serial.data_bits = databits;
	myrig->state.rigport.parm.serial.stop_bits = stopbits;
	myrig->state.rigport.parm.serial.parity = parity;
	myrig->state.rigport.parm.serial.handshake = handshake;
	strncpy(myrig->state.rigport.pathname, portname, FILPATHLEN - 1);
	printf("path: %s\n", portname);
	printf("path: %s\n", myrig->state.rigport.pathname);
	return 0;
}

int init_rig(int rig_model)
{
	//check if rig already exists
	if (myrig != 0) return -1;

	rig_load_all_backends();
	myrig = rig_init(rig_model);
	if (!myrig) {
		return -1;
	}
	
	return 0;
}

int open_rig()
{
	int res = rig_open(myrig);
	return res;
}

int set_vfo(int vfo)
{
	int res = rig_set_vfo(myrig, vfo);
	return res;
}

int set_freq(int vfo, double freq)
{
	int res = rig_set_freq(myrig, vfo, freq);
	return res;
}

int set_mode(int vfo, int mode, int pb_width)
{
	int res = rig_set_mode(myrig, vfo, mode, pb_width);
	return res;
}

int get_passband_narrow(int mode)
{
	int res = rig_passband_narrow(myrig, mode);
	return res;
}

int get_passband_normal(int mode)
{
	int res = rig_passband_normal(myrig, mode);
	return res;
}

int get_passband_wide(int mode)
{
	int res = rig_passband_wide(myrig, mode);
	return res;
}

int get_freq(int vfo, double *freq)
{
	int res = rig_get_freq(myrig, vfo, freq);
	return res;
}

int get_mode(int vfo, int *mode, int *pb_width)
{
	int res = rig_get_mode(myrig, vfo, mode, pb_width);
	return res;
}

int set_ptt(int vfo, int ptt)
{
	int res = rig_set_ptt(myrig, vfo, ptt);
	return res;
}

int get_ptt(int vfo, int *ptt)
{
	int res = rig_get_ptt(myrig, vfo, ptt);
	return res;
}

int set_rit(int vfo, int offset)
{
	int res = rig_set_rit(myrig, vfo, offset);
	return res;
}

int get_rit(int vfo, int *offset)
{
	int res = rig_get_rit(myrig, vfo, offset);
	return res;
}

int set_xit(int vfo, int offset)
{
	int res = rig_set_xit(myrig, vfo, offset);
	return res;
}

int get_xit(int vfo, int *offset)
{
	int res = rig_get_xit(myrig, vfo, offset);
	return res;
}

int set_split_freq(int vfo, double tx_freq)
{
	int res = rig_set_split_freq(myrig, vfo, tx_freq);
	return res;
}

int get_split_freq(int vfo, double *tx_freq)
{
	int res = rig_get_split_freq(myrig, vfo, tx_freq);
	return res;
}

int set_split_mode(int vfo, int tx_mode, int tx_width)
{
	int res = rig_set_split_mode(myrig, vfo, tx_mode, tx_width);
	return res;
}

int get_split_mode(int vfo, int *tx_mode, int *tx_width)
{
	int res = rig_get_split_mode(myrig, vfo, tx_mode, tx_width);
	return res;
}

int set_split_vfo(int vfo, int split, int tx_vfo)
{
	int res = rig_set_split_vfo(myrig, vfo, split, tx_vfo);
	return res;
}

int get_split_vfo(int vfo, int *split, int *tx_vfo)
{
	int res = rig_get_split_vfo(myrig, vfo, split, tx_vfo);
	return res;
}

int set_powerstat(int status)
{
	int res = rig_set_powerstat(myrig, status);
	return res;
}

int get_powerstat(int *status)
{
	int res = rig_get_powerstat(myrig, status);
	return res;
}

const char* get_info()
{
	char *info;
	info = rig_get_info(myrig);
	return info;
}

int set_ant(int vfo, int ant)
{
	int res = rig_set_ant(myrig, vfo, ant);
	return res;
}

int get_ant(int vfo, int *ant)
{
	int res = rig_get_ant(myrig, vfo, ant);
	return res;
}

int set_ts(int vfo, int ts)
{
	int res = rig_set_ts(myrig, vfo, ts);
	return res;
}

int get_ts(int vfo, int *ts)
{
	int res = rig_get_ts(myrig, vfo, ts);
	return res;
}

unsigned long has_get_level(unsigned long level)
{
	unsigned long res = rig_has_get_level(myrig, level);
	return res;
}

unsigned long has_set_level(unsigned long level)
{
	unsigned long res = rig_has_set_level(myrig, level);
	return res;
}

unsigned long has_get_func(unsigned long func)
{
	unsigned long res = rig_has_get_func(myrig, func);
	return res;
}

unsigned long has_set_func(unsigned long func)
{
	unsigned long res = rig_has_set_func(myrig, func);
	return res;
}

unsigned long has_get_parm(unsigned long parm)
{
	unsigned long res = rig_has_get_parm(myrig, parm);
	return res;
}

unsigned long has_set_parm(unsigned long parm)
{
	unsigned long res = rig_has_set_parm(myrig, parm);
	return res;
}

int get_level(int vfo, unsigned long level, int *value)
{	
	value_t v;
	int res = rig_get_level(myrig, vfo, level, &v);

	switch(level) {
		case RIG_LEVEL_AF:
		case RIG_LEVEL_RF:
		case RIG_LEVEL_NR:
		case RIG_LEVEL_RFPOWER:
		case RIG_LEVEL_MICGAIN:
		case RIG_LEVEL_SQL:
		case RIG_LEVEL_COMP:
		case RIG_LEVEL_VOXGAIN:
		case RIG_LEVEL_ANTIVOX:
		case RIG_LEVEL_SWR:
		case RIG_LEVEL_ALC:
			*value = (int) (v.f*100);
			break;
		case RIG_LEVEL_NONE:
		case RIG_LEVEL_PREAMP:
		case RIG_LEVEL_ATT:
		case RIG_LEVEL_AGC:
		case RIG_LEVEL_IF:
		case RIG_LEVEL_CWPITCH:
		case RIG_LEVEL_KEYSPD:
		case RIG_LEVEL_NOTCHF:
		case RIG_LEVEL_VOX:
		case RIG_LEVEL_BKINDL:
		case RIG_LEVEL_METER:
		case RIG_LEVEL_STRENGTH:
		case RIG_LEVEL_RAWSTR:
			*value = (int)v.i;
			break;
		default: 
			*value = 0;
			printf("Unknown Level in 'get_level_'; Conversion not possible");
			return RIG_EINVAL;
	}

	return res;
} 

int set_level(int vfo, unsigned long level, int value)
{	
	value_t v;

	switch(level) {
		case RIG_LEVEL_AF:
		case RIG_LEVEL_RF:
		case RIG_LEVEL_NR:
		case RIG_LEVEL_RFPOWER:
		case RIG_LEVEL_MICGAIN:
		case RIG_LEVEL_SQL:
		case RIG_LEVEL_COMP:
		case RIG_LEVEL_VOXGAIN:
		case RIG_LEVEL_ANTIVOX:
		case RIG_LEVEL_SWR:
		case RIG_LEVEL_ALC:
			v.f = ((float)value/100);
			printf ("power: %f", v.f);
			break;
		case RIG_LEVEL_NONE:
		case RIG_LEVEL_PREAMP:
		case RIG_LEVEL_ATT:
		case RIG_LEVEL_AGC:
		case RIG_LEVEL_IF:
		case RIG_LEVEL_CWPITCH:
		case RIG_LEVEL_KEYSPD:
		case RIG_LEVEL_NOTCHF:
		case RIG_LEVEL_VOX:
		case RIG_LEVEL_BKINDL:
		case RIG_LEVEL_METER:
		case RIG_LEVEL_STRENGTH:
		case RIG_LEVEL_RAWSTR:
			v.i = (int) value;
			break;
		default: 
			printf("Unknown Level in 'set_level'; Conversion not possible");
			return RIG_EINVAL;
	}
	
	int res = rig_set_level(myrig, vfo, level, v);
	return res;
} 


void set_debug_level(int debug_level)
{
	rig_set_debug (debug_level);
}

int close_rig()
{
	int res = rig_close(myrig);
	return res;
}

int cleanup_rig()
{
	int res = rig_cleanup(myrig);
	return res;
}

