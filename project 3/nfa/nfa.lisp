;; You may define helper functions here
(defun move (transition states input)
    (loop for x in states append( mapcar (lambda (y) ( list y (first x))) (funcall transition (first x) input) ))
)

(defun getpath(path_table  final step)
   (if (eql step 0)
   (list final)
   (if (null (gethash (list final step) path_table))
        NIL
        (let ((path (list final))) (append path (getpath path_table (gethash (list final step) path_table) (+ -1 step))))))
)

(defun reachable (transition start final input)
    ;; TODO: Incomplete function
    ;; The next line should not be in your solution.
    (let ((queue (list (list start nil))) (cnt 0))
         (setq table (make-hash-table :test 'equal))
          (setf (gethash (list start 0) table) nil) 
          (loop for label in input do (
            (lambda (label) 
                (let ((newqueue (move transition queue label))) 
                    (incf cnt)
                    (setq queue (remove-duplicates newqueue :test 'equal))
                    ;(print queue)
                    (loop for state in queue do ( (lambda(state) (let() (setf (gethash (list (first state) cnt) table) (second state)) )) state) )
                    )) label)
          )
          ;(print table)
           (reverse (getpath table final cnt)))
)