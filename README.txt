==========
freecrypto
==========

The *freecrypto* library is a set of independent cryptography modules without
copyright.

The goal is to have a cryptography library without copyright to avoid the
virality about regulations from USA. To get it is used `Creative Commons CC0`_.

Note that Go's cryptography library is copyrighted by an USA company, so it's
liable to cryptography regulations from USA. That viral law affects to any
program where it's being used, so it can not be used in some third countries
(where USA says you).

It’s far to be finished but slowly may be terminated.


.. _Creative Commons CC0: http://wiki.creativecommons.org/CC0_FAQ

Installation
============

By now, the only useful module is `rand` which allows to get random bytes::

	$ [sudo -E] goinstall github.com/kless/freecrypto/rand

