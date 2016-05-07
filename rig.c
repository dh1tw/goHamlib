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
