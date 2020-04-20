
// +build !cmd_go_bootstrap

package batsky

import _ "unsafe"

// // #include <sysdep.h> not find
// #include <sys/time.h>
// #include <errno.h>
// #include <stdio.h>
// #include <stdlib.h>
// #include <sys/types.h>
// #include <unistd.h>
// #include <sys/socket.h>
// #include <sys/un.h>
// #include <sys/stat.h> // mkdir
// #include <emmintrin.h>
// 
// void _create_and_wait_connection(void);
// void _get_batsky_time(struct timeval *tv);
// 
// #define BATSKY_SOCK_DIR "/tmp/batsky/"
// 
// int batsky_init = 0;
// 
// int batsky_server_sockfd, batsky_client_sockfd;
// socklen_t batsky_client_len;
// struct sockaddr_un batsky_server_address;
// struct sockaddr_un  batsky_client_address;
// 
// static int batsky_lock = 0;
// 
// void ___spin_lock(int volatile *p)
// {
//     while(!__sync_bool_compare_and_swap(p, 0, 1))
//     {
//         // spin read-only until a cmpxchg might succeed
//         while(*p) _mm_pause();  // or maybe do{}while(*p) to pause first
//     }
// }
// 
// void ___spin_unlock(int volatile *p)
// {
//     asm volatile ("":::"memory"); // acts as a memory barrier.
//     *p = 0;
// }
// 
// void _create_and_wait_connection(void) {
//     char batsky_sock_name[256];
//     pid_t pid = getpid();
// 
//     /* create socket */
//     snprintf(batsky_sock_name, sizeof batsky_sock_name, "%s/%d_batsky.sock", BATSKY_SOCK_DIR, pid);
//     unlink(batsky_sock_name);
//     batsky_server_sockfd = socket(AF_UNIX, SOCK_STREAM, 0);
// 
//     batsky_server_address.sun_family = AF_UNIX;
//     strcpy(batsky_server_address.sun_path, batsky_sock_name);
// 
//     int ret = bind(batsky_server_sockfd, (struct sockaddr *)&batsky_server_address, sizeof(batsky_server_address));
//     if (ret) {
//         perror("Bind for batsky_socket failed");
//     }
//     
//     listen(batsky_server_sockfd, 1);
// 
//     /*  Accept a connection.  */
//     batsky_client_len = sizeof(batsky_client_address);
//     batsky_client_sockfd = accept(batsky_server_sockfd, (struct sockaddr *)&batsky_client_address, &batsky_client_len);
// 
//     write(batsky_client_sockfd, &pid, 4);
//    
// }
// 
// void _get_batsky_time(struct timeval *tv) {
//     int n;
//     // Ask batsky
//     n = write(batsky_client_sockfd, &tv->tv_sec, 8);
//     if (n != 8) perror("Write incomplete");
//     n = write(batsky_client_sockfd, &tv->tv_usec, 8);
//     if (n != 8) perror("Write incomplete");
//     // Receive simulated time
//     n = read(batsky_client_sockfd, &tv->tv_sec, 8);
//     if (n != 8) perror("Read incomplete");
//     n = read(batsky_client_sockfd, &tv->tv_usec, 8);
//     if (n != 8) perror("Read incomplete");
//     
// }
// 
// struct timeval
// gettimeofday_batsky1 ()
// {
//
//     struct timeval tv;
//     //struct timezone tz;
//     // int ret = INLINE_SYSCALL (gettimeofday, 2, &tv, &tz); // MACRO provided sysdep.h
//     int ret = gettimeofday(&tv, NULL);
//
//     ___spin_lock(&batsky_lock);
// 
//     /* if BATSKY_SOCK_DIR does not exist return the orginal gettimeofday's result */
//     if( access( BATSKY_SOCK_DIR, F_OK ) != -1 ) {
//         if (batsky_init == 0) {
//             printf("create_and_wait_connection");
//             _create_and_wait_connection();
//             batsky_init = 1;
//         }  
//         _get_batsky_time(&tv);
//     }
//     ___spin_unlock(&batsky_lock);
//
//     return tv;
// }
import "C"

//go:nosplit
func nanotime() int64 {
	tv := C.gettimeofday_batsky1()
	return int64(tv.tv_sec) * 1e9 + int64(1000 * tv.tv_usec)
}
