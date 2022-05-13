# GoProject

# Definition

1>Orchestrator
  Automates deploying, scaling & managing containers.
2>Task
  Smallest unit of work.
    Should specify it's needs for memory, CPU & disk.
    What to do in case of failure (restart_policy).
    Name of container image used to run the task.
3>Job (not included in this implementation)
  Group of tasks to perform a set of functions.
  Should specify details at a higher level than tasks.
    Each task that makes up the job.
    Which data centers the job should run in.
    The type of job.
4>Scheduler
  Decides which machine can best host the tasks defined.
  Could be a process within the machine or possibly a different service.
  Decisions can vary in model complexity from round-robin to score calculation based on variables and deciding node by best score
    Detrmine set of candidate machines to run task
    Score machines from best to worst
    Pick machine with best score
5>Manager
  Brain of orchestrator & main entry point
  Jobs go to manager, manager uses scheduler to determine where to run job.
  Can gather metrics, to use for scheduling.
    Accept requests from users to start/stop tasks
    Schedule tasks onto worker machines
    Keep track of tasks, states & the machine they run on
6>Worker
  Runs the tasks assigned to it, attempts task restarts on failure, makes metrics about task and machine health for master
    Runs asks as containers
    Accepts tasks to run from manager
    provides relevant statistics to the manager to schedule tasks
    Keeps track of tasks & their state
7>Cluster
  Logical grouping of all above components.
8>CLI
  Main user interface
    Start/Stop Tasks
    Get status of tasks
    See state of machines
    Start manager
    Start worker
