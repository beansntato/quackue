# quackue

A job queue that guarantees every job runs exactly once, even when the process crashes at any point during the job lifecycle.

Distributed systems can fail mid-operation. Without careful design, a crash either loses the job (it never ran) or doubles it (it runs again on recovery). This system proves correctness across both failure modes simultaneously.

This project will make use of the [walrus](https://github.com/beansntato/walrus) for the Write-Ahead Log (WAL), and a Postgres Job Queue. Every state transition is ordered, so recovery can always reconstruct what happened and safely replay. The walrus contains a State Machine interface that the project can implement. If needed, walrus project can be updated.
