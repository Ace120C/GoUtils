## Understanding the code:

`os/user` is a go package that specializes in finding out users and groups, using `user.Current()`
will display your gid, username, and your filepath, this function expects an error so you must error handle it 
accordingly.
