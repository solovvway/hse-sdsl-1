.section .data
msg1:
    .ascii "Hello, World!\n"
    len = . - msg1         
msg2:
    .ascii "Hello, World!\n"
    len = . - msg2        
.section .text
.globl _start             