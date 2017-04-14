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
	// printf("path: %s\n", portname);
	// printf("path: %s\n", myrig->state.rigport.pathname);
	return 0;
}

int init_rig(int rig_model)
{
	//check if rig already exists
	if (myrig != 0) return -1;

	// rig_load_all_backends();
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

int has_set_vfo()
{
	if (myrig->caps->set_vfo)
	{
		return RIG_OK;
	}
	return RIG_ENIMPL;
}

int set_vfo(int vfo)
{
	int res = rig_set_vfo(myrig, vfo);
	return res;
}

int has_get_vfo()
{
	if (myrig->caps->get_vfo)
	{
		return RIG_OK;
	}
	return RIG_ENIMPL;
}

int get_vfo(int *vfo)
{
	int res = rig_get_vfo(myrig, vfo);
	return res;
}

int has_set_freq()
{
	if (myrig->caps->set_freq)
	{
		return RIG_OK;
	}
	return RIG_ENIMPL;
}

int set_freq(int vfo, double freq)
{
	int res = rig_set_freq(myrig, vfo, freq);
	return res;
}

int has_set_mode()
{
	if (myrig->caps->set_mode)
	{
		return RIG_OK;
	}
	return RIG_ENIMPL;
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

int has_get_freq()
{
	if (myrig->caps->get_freq)
	{
		return RIG_OK;
	}
	return RIG_ENIMPL;
}

int get_freq(int vfo, double *freq)
{
	int res = rig_get_freq(myrig, vfo, freq);
	return res;
}

int has_get_mode()
{
	if (myrig->caps->get_mode)
	{
		return RIG_OK;
	}
	return RIG_ENIMPL;
}

int get_mode(int vfo, int *mode, long *pb_width)
{
	rmode_t* m = (rmode_t*)mode;
	pbwidth_t* pb = (pbwidth_t*)pb_width;

	int res = rig_get_mode(myrig, vfo, m, pb);
	return res;
}

int has_set_ptt()
{
	if (myrig->caps->set_ptt)
	{
		return RIG_OK;
	}
	return RIG_ENIMPL;
}

int set_ptt(int vfo, int ptt)
{
	int res = rig_set_ptt(myrig, vfo, ptt);
	return res;
}

int has_get_ptt()
{
	if (myrig->caps->get_ptt)
	{
		return RIG_OK;
	}
	return RIG_ENIMPL;
}

int get_ptt(int vfo, int *ptt)
{
	int res = rig_get_ptt(myrig, (vfo_t)vfo, (ptt_t*)ptt);
	return res;
}

int has_set_rit()
{
	if (myrig->caps->set_rit)
	{
		return RIG_OK;
	}
	return RIG_ENIMPL;
}

int set_rit(int vfo, int offset)
{
	int res = rig_set_rit(myrig, vfo, offset);
	return res;
}

int has_get_rit()
{
	if (myrig->caps->get_rit)
	{
		return RIG_OK;
	}
	return RIG_ENIMPL;
}

int get_rit(int vfo, long *offset)
{
	int res = rig_get_rit(myrig, vfo, offset);
	return res;
}

int has_set_xit()
{
	if (myrig->caps->set_xit)
	{
		return RIG_OK;
	}
	return RIG_ENIMPL;
}

int set_xit(int vfo, int offset)
{
	int res = rig_set_xit(myrig, vfo, offset);
	return res;
}

int has_get_xit()
{
	if (myrig->caps->get_xit)
	{
		return RIG_OK;
	}
	return RIG_ENIMPL;
}

int get_xit(int vfo, long *offset)
{
	int res = rig_get_xit(myrig, (vfo_t)vfo, offset);
	return res;
}

int has_set_split_freq()
{
	if (myrig->caps->set_split_freq)
	{
		return RIG_OK;
	}
	return RIG_ENIMPL;
}

int set_split_freq(int vfo, double tx_freq)
{
	int res = rig_set_split_freq(myrig, vfo, tx_freq);
	return res;
}

int has_get_split_freq()
{
	if (myrig->caps->get_split_freq)
	{
		return RIG_OK;
	}
	return RIG_ENIMPL;
}

int get_split_freq(int vfo, double *tx_freq)
{
	int res = rig_get_split_freq(myrig, vfo, tx_freq);
	return res;
}

int has_set_split_mode()
{
	if (myrig->caps->set_split_mode)
	{
		return RIG_OK;
	}
	return RIG_ENIMPL;
}

int set_split_mode(int vfo, int tx_mode, int tx_width)
{
	int res = rig_set_split_mode(myrig, vfo, tx_mode, tx_width);
	return res;
}

int has_get_split_mode()
{
	if (myrig->caps->get_split_mode)
	{
		return RIG_OK;
	}
	return RIG_ENIMPL;
}

int get_split_mode(int vfo, int *tx_mode, long *tx_width)
{
	rmode_t* mode = (rmode_t*)tx_mode;
	int res = rig_get_split_mode(myrig, vfo, mode, tx_width);
	return res;
}

int has_set_split_vfo()
{
	if (myrig->caps->set_split_vfo)
	{
		return RIG_OK;
	}
	return RIG_ENIMPL;
}

int set_split_vfo(int vfo, int split, int tx_vfo)
{
	int res = rig_set_split_vfo(myrig, vfo, split, tx_vfo);
	return res;
}

int has_get_split_vfo()
{
	if (myrig->caps->get_split_vfo)
	{
		return RIG_OK;
	}
	return RIG_ENIMPL;
}

int get_split_vfo(int vfo, int *split, int *tx_vfo)
{
	split_t* sp = (split_t*)split;
	int res = rig_get_split_vfo(myrig, vfo, sp, tx_vfo);
	return res;
}

int has_set_powerstat()
{
	if (myrig->caps->set_powerstat)
	{
		return RIG_OK;
	}
	return RIG_ENIMPL;
}

int set_powerstat(int status)
{
	int res = rig_set_powerstat(myrig, status);
	return res;
}

int has_get_powerstat()
{
	if (myrig->caps->get_powerstat)
	{
		return RIG_OK;
	}
	return RIG_ENIMPL;
}

int get_powerstat(int *status)
{
	int res = rig_get_powerstat(myrig, (powerstat_t*)status);
	return res;
}

const char* get_info()
{
	const char *info;
	info = rig_get_info(myrig);
	return info;
}

void get_model_name(char *rig_name)
{
	strncpy(rig_name, myrig->caps->model_name, sizeof(myrig->caps->model_name));
}

void get_version(char *version)
{
	strncpy(version, myrig->caps->version, sizeof(myrig->caps->version));
}

void get_mfg_name(char *mfg_name)
{
	strncpy(mfg_name, myrig->caps->mfg_name, sizeof(myrig->caps->mfg_name));
}

int get_status()
{
	return myrig->caps->status;
}

int has_set_ant()
{
	if (myrig->caps->set_ant)
	{
		return RIG_OK;
	}
	return RIG_ENIMPL;
}

int set_ant(int vfo, int ant)
{
	int res = rig_set_ant(myrig, vfo, ant);
	return res;
}

int has_get_ant()
{
	if (myrig->caps->get_ant)
	{
		return RIG_OK;
	}
	return RIG_ENIMPL;
}

int get_ant(int vfo, int *ant)
{
	int res = rig_get_ant(myrig, vfo, ant);
	return res;
}

int has_set_ts()
{
	if (myrig->caps->set_ts)
	{
		return RIG_OK;
	}
	return RIG_ENIMPL;
}

int set_ts(int vfo, int ts)
{
	int res = rig_set_ts(myrig, vfo, ts);
	return res;
}

int has_get_ts()
{
	if (myrig->caps->get_ts)
	{
		return RIG_OK;
	}
	return RIG_ENIMPL;
}

int get_ts(int vfo, long *ts)
{
	int res = rig_get_ts(myrig, vfo, ts);
	return res;
}

signed long get_rig_resolution(int mode)
{
	signed long resolution = rig_get_resolution(myrig, mode);
	return resolution; 
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

int has_token(char *token)
{
	token_t t = rig_token_lookup(myrig, token);

	if (t == RIG_CONF_END)
	{
		return RIG_EINVAL;
	}
	return RIG_OK;
}

int has_get_conf()
{
	if (myrig->caps->get_conf)
	{
		return RIG_OK;
	}
	return RIG_ENIMPL;
}

int get_conf(char* token, char* val)
{
	token_t t = rig_token_lookup(myrig, token);

	if (t == RIG_CONF_END){
		return -RIG_EINVAL;
	}
	
	int res = rig_get_conf(myrig, t, val);
	return res; 
}

int has_set_conf()
{
	if (myrig->caps->set_conf)
	{
		return RIG_OK;
	}
	return RIG_ENIMPL;
}

int set_conf(char* token, char* val)
{
	token_t t = rig_token_lookup(myrig, token);

	if (t == RIG_CONF_END){
		return -RIG_EINVAL;
	}
	
	int res = rig_set_conf(myrig, t, val);
	return res; 
}

int get_level(int vfo, unsigned long level, float *value)
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
		case RIG_LEVEL_BALANCE:
		case RIG_LEVEL_PBT_IN:
		case RIG_LEVEL_PBT_OUT:
		case RIG_LEVEL_APF:		
			*value = v.f;
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
		case RIG_LEVEL_SLOPE_LOW:
		case RIG_LEVEL_SLOPE_HIGH:
		case RIG_LEVEL_BKIN_DLYMS:
			*value = (float)v.i;
			break;
		default: 
			*value = 0;
			printf("Unknown Level in 'get_level'; Conversion not possible\n");
			return RIG_EINVAL;
	}

	return res;
} 

int get_level_gran(unsigned long level, float *step, float *min, float *max)
{
	gran_t gran;
	int found = 0;
	int i;

	for (i = 0; i < RIG_SETTING_MAX; i++) {

		if (level & rig_idx2setting(i))
		{
			gran = myrig->caps->level_gran[i];	
			found = 1;
			break;
		}
	}

	//return if no gran was found for this level
	if (found != 1)
	{
		return RIG_EINVAL;
	}

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
		case RIG_LEVEL_BALANCE:
		case RIG_LEVEL_PBT_IN:
		case RIG_LEVEL_PBT_OUT:
		case RIG_LEVEL_APF:
			*min = gran.min.f;
			*max = gran.max.f;
			*step = gran.step.f;
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
		case RIG_LEVEL_SLOPE_LOW:
		case RIG_LEVEL_SLOPE_HIGH:
		case RIG_LEVEL_BKIN_DLYMS:
			*min = (float)gran.min.i;
			*max = (float)gran.max.i;
			*step = (float)gran.step.i;
			break;
		default: 
			*min = 0;
			*max = 0;
			*step = 0;
			printf("Unknown Level in 'get_level_gran'; Conversion not possible\n");
			return RIG_EINVAL;
	}
	return RIG_OK;
}


int set_level(int vfo, unsigned long level, float value)
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
		case RIG_LEVEL_BALANCE:
		case RIG_LEVEL_PBT_IN:
		case RIG_LEVEL_PBT_OUT:
			v.f = value;
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
		case RIG_LEVEL_SLOPE_LOW:
		case RIG_LEVEL_SLOPE_HIGH:
			v.i = (int) value;
			break;
		default: 
			printf("Unknown Level in 'set_level'; Conversion not possible\n");
			return RIG_EINVAL;
	}
	
	int res = rig_set_level(myrig, vfo, level, v);
	return res;
}
 
int get_func(int vfo, unsigned long function, int *value)
{
	int res = rig_get_func(myrig, vfo, function, value);
	return res;
} 

int set_func(int vfo, unsigned long function, int value)
{
	int res = rig_set_func(myrig, vfo, function, value);
	return res;
}

//couldn't find an example to check this function properly
//it might be necessary to cast return values to float or char*
int get_parm(unsigned long parm, float *value)
{	
	value_t v;
	int res = rig_get_parm(myrig, parm, &v);

	switch(parm) {
		case RIG_PARM_NONE:
		case RIG_PARM_ANN:
		case RIG_PARM_APO:
		case RIG_PARM_BACKLIGHT:
		case RIG_PARM_BEEP:
		case RIG_PARM_TIME:
		case RIG_PARM_BAT:
		case RIG_PARM_KEYLIGHT:
			*value = (float)v.i;
			break;
		default: 
			*value = 0;
			printf("Unknown Parameter in 'get_param'; Conversion not possible\n");
			return RIG_EINVAL;
	}

	return res;
} 

int get_parm_gran(unsigned long parm, float *step, float *min, float *max)
{
	gran_t gran;
	int found = 0;
	int i;

	for (i = 0; i < RIG_SETTING_MAX; i++) {

		if (parm & rig_idx2setting(i))
		{
			gran = myrig->caps->parm_gran[i];	
			found = 1;
			break;
		}
	}

	//return if no gran was found for this level
	if (found != 1)
	{
		return RIG_EINVAL;
	}

	switch(parm) {
		case RIG_PARM_NONE:
		case RIG_PARM_ANN:
		case RIG_PARM_APO:
		case RIG_PARM_BACKLIGHT:
		case RIG_PARM_BEEP:
		case RIG_PARM_TIME:
		case RIG_PARM_BAT:
		case RIG_PARM_KEYLIGHT:
			*min = (float)gran.min.i;
			*max = (float)gran.max.i;
			*step = (float)gran.step.i;
			break;
		default: 
			*min = 0;
			*max = 0;
			*step = 0;
			printf("Unknown Parameter in 'get_parm_gran'; Conversion not possible\n");
			return RIG_EINVAL;
	}
	return RIG_OK;
}

//couldn't find an example to check this function properly
//it might be necessary to cast return values to float or char*
int set_parm(unsigned long parm, float value)
{	
	value_t v;

	switch(parm) {
		case RIG_PARM_NONE:
		case RIG_PARM_ANN:
		case RIG_PARM_APO:
		case RIG_PARM_BACKLIGHT:
		case RIG_PARM_BEEP:
		case RIG_PARM_TIME:
		case RIG_PARM_BAT:
		case RIG_PARM_KEYLIGHT:
			v.i = (int) value;
			break;
		default: 
			printf("Unknown Parameter in 'set_param'; Conversion not possible\n");
			return RIG_EINVAL;
	}
	
	int res = rig_set_parm(myrig, parm, v);
	return res;
}

int vfo_op(int vfo, int op)
{
	int res;
	res = rig_vfo_op(myrig, (vfo_t)vfo, (vfo_op_t)op);
	return res;
}

int get_caps_max_rit(int *rit)
{
	*rit = myrig->caps->max_rit;
	return RIG_OK;
}

int get_caps_max_xit(int *xit)
{
	*xit = myrig->caps->max_xit;
	return RIG_OK;
}

int get_caps_max_if_shift(int *if_shift)
{
	*if_shift = myrig->caps->max_ifshift;
	return RIG_OK;
}

int* get_caps_attenuator_list_pointer_and_length(int *length)
{
	*length = sizeof(myrig->caps->attenuator)/sizeof(int);
	return myrig->caps->attenuator;
}

int* get_caps_preamp_list_pointer_and_length(int *length)
{
	*length = sizeof(myrig->caps->preamp)/sizeof(int);
	return myrig->caps->preamp;
}

int get_supported_vfos(int *vfo_list)
{
	*vfo_list = (int)myrig->state.vfo_list;
	return RIG_OK;
}

int get_supported_vfo_operations(int *vfo_ops)
{
	*vfo_ops = (int)myrig->caps->vfo_ops;
	return RIG_OK;
}

int get_supported_modes(int *modes)
{
	*modes = (int)myrig->state.mode_list;
	return RIG_OK;
}

int get_filter_count(int *filter_count)
{
	int i;
	for (i=0; i<FLTLSTSIZ && !RIG_IS_FLT_END(myrig->caps->filters[i]); i++)
	{
		*filter_count += 1;
	}
	return RIG_OK;
}

int get_filter_mode_width(int filter, int *mode, signed long *width)
{
	*mode = myrig->caps->filters[filter].modes;	
	*width = myrig->caps->filters[filter].width;

	return RIG_OK;
} 

int get_ts_count(int *ts_count)
{
	int i;
	*ts_count = sizeof(myrig->caps->tuning_steps)/ sizeof(myrig->caps->tuning_steps[0]);
	return RIG_OK;
}

int get_tuning_steps(int el, int *mode, signed long *ts)
{
	*mode = myrig->caps->tuning_steps[el].modes;
	*ts = myrig->caps->tuning_steps[el].ts;

	return RIG_OK;
}

int get_int_from_array(int *array, int *el, int index)
{
	*el = array[index];
	return RIG_OK;
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
	myrig = 0;
	return res;
}

