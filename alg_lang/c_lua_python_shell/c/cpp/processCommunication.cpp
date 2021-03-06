
/* 近期在写个程序的时候须要在进程间通讯，详细需求是这样。

1.       主要有两个进程：一个进程作为被请求进程，我们称为 SERVER 进程；还有一个进程是请求进程，称为 CLIENG 进程。

2.       SERVER 进程提供一些服务，其完毕计算功能；而 CLIENT 进程须要在它运行完计算之后将结果取会。

 

因为计算结果可能是一个结构，也可能是一个复杂的数据，所以通过消息来在进程传递信息是有限的。还有一方面通常是单方向的通讯，实际上这里的需求有一个双向性，看下图：
ProcessComm.JPG


这里两个进程都能够有自己的窗体，因此实际上我们能够通过消息来通知对方。但细致一想，请求服务通过windows消息是没有问题的，通知结果通过消息是不妥当的，实际上我们须要在请求服务完以后马上得到运行结果，而使用windows消息可能有时间上的问题，并且同步很麻烦。

想要的结果是在 CLIENT 请求服务以后马上得到返回结果，但我们能够改变一下思路就easy多了：结果不是有 SERVER 返回，而由 CLIENG 自己去获取。这样，我们能够在请求消息的时候使用 SendMessage 来发送这个请求。这是必须的， SendMessage 发送消息是同步的方式，必须等到消息处理完成之后才返回，而这正是我们想要的结果。那么 CLIENT 怎么样才干获得 SERVER 进程的结果来。

我们知道， WINDOWS 下每一个进程都有自己的地址空间，普通情况下一个进程訪问你一个进程的地址是不对的，最简单的是提示该地址无效，程序崩溃。当然还是有非常多中方式来读取对方进程的地址空间上的数据。

1．  能够做一个 DLL ，利用注入 DLL 的方式，来将该 DLL 注入到远程进程的地址空间上，然后通过 DLL 的 API 来读取。这个方式有点麻烦，还必须写一个 DLL ，还要注入 DLL 。

2．  使用 WriteProcessMemory 和 ReadProcessMemory 来读写远程进程的内存。这样的方式相对照较简单，其有一个參数远程进程的 HANDLE 指示你须要读写哪个进程的内存。当然，使用这两个函数是须要远程进程的地址空间的，这个地址是通过 VirtualAllocEx 来分配的，其通过 VirtualFreeEx 来释放。这两个 API 也有一个进程句柄參数。

有人问，为什么不使用 WM_COPYDATA 来传递大块数据？好，该消息是能够传递大块数据，但我们的应用中须要消息的双向，所以这里使用 WM_COPYDATA 是不合适的。

 

以下我们看一下第 2 种方法的详细步骤：

1．  找到对方进程的窗体。

2．  找到他的进程 ID

3．  使用 PROCESS_VM_OPERATION| PROCESS_VM_WRITE|PROCESS_VM_READ 标志来打开该进程，得到该继承的 HANDLE 。

4．  使用 VirtualAllocEx 来在该进程上分配适当大小的内存，得到一个地址，这个地址是远程进程的，通过不同的方式来改动该地址上的值是无效的。

5．  假设须要传递一些參数到远程进程，我们能够在该内存上写一些内容，通过 WriteProcessMemory 来完毕

6．  使用 SendMessage 来发送请求服务消息，同一时候将上面分配的内存地址作为參数传递给远程进程。

7．  远程进程得到指定消息后处理该消息，取得參数，计算结果，将结果写到指定的地址。因为这个地址是远程进程自己的地址空间，其操作这块内存的方法没有什么特别之处。

8．  SendMessage 消息返回， CLIENT 知道后，从上面的内存地址中读取返回结果；这里必须使用 ReadProcessMemory 来读取。好了，整个过程结束。

9．  调用 VritualFreeEx 将上面的内存释放。

这里须要强调两点：

1．  打开进程的时候必须设置对虚拟内存可操作、可写、可读，假设仅仅是可写，那么 ReadProcessMemory 将读取不对。

2．  必须释放该内存。