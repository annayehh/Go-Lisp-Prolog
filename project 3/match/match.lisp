; You may define helper functions here

(defun match (pattern assertion)
  (cond
   ;;Match if both empty
   ((and (null pattern) (null assertion)) 'T)
   ;;Doesn't match if only one empty
   ((or (null pattern) (null assertion)) nil)
   ;;? case
   ((eql (car pattern) '?)
    (and assertion (match (cdr pattern) (cdr assertion))))
   ;;! case
   ((eql (car pattern) '!)
     (or (and (or (eql (car (cdr pattern)) '!)
                  (eql (car (cdr pattern)) '?)
                  (eql (car (cdr pattern)) (car (cdr assertion))))
              (match (cdr pattern) (cdr assertion)))
         (and assertion (match pattern (cdr assertion)))))
   ;;car equal case
   ((eql (car pattern) (car assertion))
    (match (cdr pattern) (cdr assertion)))
   ;;Otherwise
   (t nil)
 )
)
