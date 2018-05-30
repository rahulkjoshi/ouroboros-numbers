import sys

from absl import app
from absl import flags
from absl import logging

from typing import List, Sequence, TypeVar, Iterable, Optional

T = TypeVar('T')

FLAGS = flags.FLAGS

flags.DEFINE_integer('num_digits', 2, 'Length of the number.')
flags.DEFINE_integer('multiplier', 6, 'Number to multiply by.')

def digits() -> Iterable[int]:
    return range(0, 10)

def letters(length: int) -> Iterable[str]:
    if length > 26:
        return []
    for l in range(ord('A'), ord('A') + length):
        yield chr(l)

def rotate(input_list: List[T]) -> List[T]:
    return input_list[-1:] + input_list[:-1]

def concat_ints(input_list: Iterable[int]) -> int:
    ret_int = 0
    for i in input_list:
        ret_int = ret_int * 10 + i
    return ret_int

def verify(answer_list: List[int]) -> bool:
    answer = concat_ints(answer_list)
    result_list = rotate(answer_list)
    result = concat_ints(result_list)

    return (answer * FLAGS.multiplier) == result

def print_answer(answer_list: List[int]) -> None:
    answer = concat_ints(answer_list)
    result_list = rotate(answer_list)
    result = concat_ints(result_list)

    print('Answer is: ' + str(answer))
    print('{answer} * {multiplier} = {result}'.format(
        answer=answer, multiplier=FLAGS.multiplier,
        result=(answer * FLAGS.multiplier)))
    sys.exit(0)

def main(argv):
    del argv  # Unused.
    
    if FLAGS.num_digits == 1:
        if FLAGS.multiplier == 1:
            print_answer([1])
        else:
            print('No answer possible')
        return

    number_alpha = [l for l in letters(FLAGS.num_digits)]
    result_alpha = rotate(number_alpha)

    print('Solving for:')
    print('  ' + ''.join(number_alpha))
    print('x ' + '{num:>{width}}'.format(num=FLAGS.multiplier,
                                         width=FLAGS.num_digits))
    print('-'*(FLAGS.num_digits + 2))
    print('  ' + ''.join(result_alpha))

    print('Possible ones digits:')
    ones = set([(FLAGS.multiplier * i % 10) for i in digits()])
    print(ones)

if __name__ == '__main__':
    app.run(main)
