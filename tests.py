import os, sys
import logging
logging.basicConfig(level=logging.INFO, format='%(asctime)s %(message)s', datefmt='%m/%d/%Y %I:%M:%S %p')
import click
import random
import statistics
import subprocess
import time

groups = {"G0": ["F", "R", "U", "B", "L", "D", "F'", "R'", "U'", "B'", "L'", "D'"],
            "G1": ["R", "U", "L", "D", "R'", "U'", "L'", "D'","F2", "B2", "F'2", "B'2"],
            "G2": ["U", "D", "U'", "D'","F2", "R2", "B2", "L2", "F'2", "R'2", "B'2", "L'2"],
            "G3": ["F2", "R2", "U2", "B2", "L2", "D2", "F'2", "R'2", "U'2", "B'2", "L'2", "D'2"]}

def scramble(group):
    steps_idx = []
    for _ in range(20):
        rand = random.randint(0, len(group)-1)
        while len(steps_idx) and group[rand][0] == group[steps_idx[-1]][0]:
            rand = random.randint(0, len(group)-1)
        steps_idx.append(rand)

    steps = " ".join([group[idx] for idx in steps_idx])
    print("Applying", steps, ":")
    return steps

@click.command()
@click.argument("group", default="G0")
def main(group):
    counts = []
    durations = []
    errors = 0
    success = 0
    # Compile solver
    if not os.path.exists("Rubik") or not os.path.exists("Rubik.exe"):
        args = ("go", "build")
        popen = subprocess.Popen(args)
    try:
        for _ in range(100):
            mix = scramble(groups[group])
            logging.info("Solving...")
            if sys.platform == "win32":
                args = ("./Rubik.exe", mix)
            else:
                args = ("./Rubik", mix)
            popen = subprocess.Popen(args, stdout=subprocess.PIPE)
            try:
                start = time.time()
                popen.wait(2*60)
                end = time.time()
                durations.append(end-start)
                output = popen.stdout.read().decode()
            except subprocess.TimeoutExpired:
                output = ""
            if len(output):
                res = str(output).replace("\n", "")
                if len(res):
                    count = len(res.split(" "))
                    counts.append(count)
                    logging.info("Done in %d steps\n" % count)
                    success += 1
                else:
                    time.sleep(0.1)
            else:
                logging.info("Error : Timeout\n")
                errors += 1
    except KeyboardInterrupt:
        print("\nAverage solution length: %d" % statistics.mean(counts))
        print("Average compute time: %2.2fs" % statistics.mean(durations))
        print("%d Success and %d Timeouts, rate = %.2f" % (success, errors, success/(success + errors)))

if __name__ == "__main__":
    main()