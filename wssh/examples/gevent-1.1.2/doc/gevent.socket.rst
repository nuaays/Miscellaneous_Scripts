====================================================================
 :mod:`gevent.socket` -- Cooperative low-level networking interface
====================================================================

This module provides socket operations and some related functions. The
API of the functions and classes matches the API of the corresponding
items in the standard :mod:`socket` module exactly, but the
synchronous functions in this module only block the current greenlet
and let the others run.

.. tip:: gevent's sockets, like most gevent objects, have thread
         affinity. That is, they can only be used from the operating
         system thread that created them (any greenlet in that thread
         can use the socket). The results of attempting to use the
         socket in another thread (for example, passing it to the
         threadpool) are not defined (but one common outcome is a
         :exc:`~gevent.hub.LoopExit` exception).

For convenience, exceptions (like :class:`error <socket.error>` and
:class:`timeout <socket.timeout>`) as well as the constants from the
:mod:`socket` module are imported into this module. In almost all
cases one can simply replace ``import socket`` with ``from gevent
import socket`` to start using cooperative sockets with no other
changes (or use :func:`gevent.monkey.patch_socket` at startup if code
changes are not desired or possible).

Standard Library Interface
==========================

The exact API exposed by this module varies depending on what version
of Python you are using. The documents below describe the API for
Python 2 and Python 3, respectively.

.. note:: All the described APIs should be imported from
   ``gevent.socket``, and *not* from their implementation modules.
   Their organization is an implementation detail that may change at
   any time.

.. toctree::

   Python 3 interface <gevent._socket3>
   Python 2 interface <gevent._socket2>

Gevent Extensions
=================

Beyond the basic standard library interface, ``gevent.socket``
provides some extensions. These are identical and shared by all
versions of Python.

Waiting
-------

These functions are used to block the current greenlet until an open
file (socket) is ready to perform I/O operations. These are low-level
functions not commonly used by many programs.

.. note:: These use the underlying libev ``io`` watchers, which means
          that they share the same implementation limits. For example,
          on some platforms they can be used with more than just
          sockets, while on others the applicability is more limited
          (POSIX platforms like Linux and OS X can use pipes and fifos
          but Windows is limited to sockets).

.. autofunction:: gevent.socket.wait_read
.. autofunction:: gevent.socket.wait_write
.. autofunction:: gevent.socket.wait_readwrite
.. autofunction:: gevent.socket.wait
.. autofunction:: gevent.socket.cancel_wait
