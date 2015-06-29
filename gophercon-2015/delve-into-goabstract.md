# Delve into Go

# Abstract

For any software engineer, a good debugger is an essential tool. With Go lacking a proper debugger, I decided to write Delve, a debugger specifically tailed towards Go. There are enough fundamental differences with Go as a language that make using a traditional debugger such as GDB unfeasible. Among those differences is the runtime scheduler.

In this talk I would like to present what makes Delve different, how it handles Go's scheduler and multi-threaded nature, and hopefully help demystify debuggers along the way.

So just what _is_ involved in debugging running software? How can an outsider glimpse into the world of a running process? When your software is behaving poorly, how can you glean insight into the reasons behind its unruliness?

To answer these questions we will take a deep dive into Delve. We will explore what makes software like Delve work, how it aims to solve problems with existing debuggers with regard to Go, and how to leverage Delve to gain insight into the behavior of your software.

This talk will be beneficial to all levels of Go developers. Discussions around the usage of Delve to debug software will appeal to all experience levels, and the breakdown of the inner workings of Delve should appeal to intermediate and advanced developers.

# Details

The theme of the talk will be split in half, diving deep in, building a strong foundation for the higher level aspects discussed later.

In the first half of the talk I'd like to introduce Delve, and possibly show a brief demonstration of it to pique interest, and then dive into all of the nitty gritty details of how Delve was able to do what was just demonstrated. We will learn how Delve came to be, how it works behind the scenes, and why it's different from existing debuggers.

The second half of the talk will cover more high level topics. Such topics will include usage of Delve, a thanks to the community for contributions and a road map for the future. The talk would end with a demonstration of Delve showing off some specific features.

# Pitch

A good debugger is something that has been missing from the Go ecosystem, and a lot of developers have felt that pain. Developers coming from Ruby have pry, and developers coming from C/C++ have GDB (which is unreliable when used with Go).

I believe this talk is pertinent because it helps demystify the work of a debugger, and shows off a more viable alternative to GDB.

My qualifications for this talk include being the author of Delve, and thus having a very intimate knowledge of the problems and solutions concerning debugging software written in Go.
